import { requestApi, requestApiNoData } from '$lib/core/http/request';
import {
	ImageBaseInfoResponseSchema,
	ListImageBaseInfoSchema,
	type ImageBaseInfo,
	type RemoveIdMessage
} from './types';

export function getImageList() {
	return requestApi({ method: 'get', url: '/rest/v1/getImageList' }, ListImageBaseInfoSchema, {
		errorTitle: 'Failed to load image list'
	});
}

export async function getImageById(payload: { id: string }): Promise<ImageBaseInfo> {
	const resp = await requestApi(
		{ method: 'post', url: '/rest/v1/getImageById', data: payload },
		ImageBaseInfoResponseSchema,
		{ errorTitle: 'Failed to load image' }
	);
	return resp.item;
}

export function removeImage(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeImage', data: payload },
		{
			errorTitle: 'Failed to delete image'
		}
	);
}
