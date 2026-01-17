export type ErrorDetails = Record<string, unknown>;

export class HttpError extends Error {
	name = 'HttpError';
	status: number;
	url?: string;
	method?: string;
	data?: unknown;

	constructor(
		message: string,
		opts: { status: number; url?: string; method?: string; data?: unknown }
	) {
		super(message);
		this.status = opts.status;
		this.url = opts.url;
		this.method = opts.method;
		this.data = opts.data;
	}
}

export class ApiError extends Error {
	name = 'ApiError';
	code: string;
	constructor(message: string, opts: { code: string }) {
		super(message);
		this.code = opts.code;
	}
}

export class ParseError extends Error {
	name = 'ParseError';
}

export class SchemaError extends Error {
	name = 'SchemaError';
}
