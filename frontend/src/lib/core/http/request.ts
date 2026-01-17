import type { AxiosRequestConfig } from 'axios';
import axios from 'axios';
import { z } from 'zod';
import { ApiError, HttpError, ParseError, SchemaError } from '../errors';
import { notifyError } from '$lib/app/notifications/store';
import { api } from './client';

const APIErrorSchema = z.object({
	code: z.string(),
	message: z.string()
});

export const createApiResponseSchema = <T extends z.ZodTypeAny>(dataSchema: T) =>
	z.object({
		success: z.boolean(),
		error: APIErrorSchema.optional(),
		data: dataSchema.optional()
	});

export async function requestApi<T>(
	config: AxiosRequestConfig,
	dataSchema: z.ZodType<T>,
	opts?: { errorTitle?: string; allowNoData?: boolean }
): Promise<T> {
	try {
		const res = await api.request(config);

		const parsed = createApiResponseSchema(dataSchema).safeParse(res.data);
		if (!parsed.success) {
			throw new SchemaError('Invalid server response schema');
		}

		if (!parsed.data.success) {
			const code = parsed.data.error?.code ?? 'unknown';
			const msg = parsed.data.error?.message ?? 'Request failed';
			throw new ApiError(msg, { code });
		}

		if (parsed.data.data === undefined && !opts?.allowNoData) {
			throw new SchemaError('Server response does not contain data');
		}

		return parsed.data.data as T;
	} catch (err: unknown) {
		const normalized = normalizeRequestError(err, config);
		notifyError(normalized, opts?.errorTitle);
		throw normalized;
	}
}

export async function requestApiNoData(
	config: AxiosRequestConfig,
	opts?: { errorTitle?: string }
): Promise<void> {
	await requestApi(config, z.unknown(), { ...opts, allowNoData: true });
}

function normalizeRequestError(err: unknown, config: AxiosRequestConfig): Error {
	if (
		err instanceof ApiError ||
		err instanceof HttpError ||
		err instanceof SchemaError ||
		err instanceof ParseError
	) {
		return err;
	}

	if (err instanceof z.ZodError) {
		return new SchemaError('Invalid server response data');
	}

	if (axios.isAxiosError(err)) {
		const method = String(config.method ?? 'GET').toUpperCase();
		const url = typeof config.url === 'string' ? config.url : undefined;

		if (err.response) {
			const status = err.response.status;
			return new HttpError(`HTTP ${status} ${method} ${url ?? ''}`.trim(), {
				status,
				url,
				method,
				data: err.response.data
			});
		}

		if (err.request) {
			return new HttpError(`Network error ${method} ${url ?? ''}`.trim(), {
				status: 0,
				url,
				method
			});
		}

		return new ParseError(err.message);
	}

	return err instanceof Error ? err : new Error('Unknown error');
}
