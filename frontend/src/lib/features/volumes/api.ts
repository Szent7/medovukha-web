import { requestApi, requestApiNoData } from '$lib/core/http/request';
import {
	ListVolumeBaseInfoSchema,
	VolumeBaseInfoResponseSchema,
	type ListVolumeBaseInfo,
	type RemoveIdMessage,
	type VolumeBaseInfo
} from './types';

export function getVolumeList(): Promise<ListVolumeBaseInfo> {
	return requestApi({ method: 'get', url: '/rest/v1/getVolumeList' }, ListVolumeBaseInfoSchema, {
		errorTitle: 'Failed to load volume list'
	});
}

export async function getVolumeById(payload: { id: string }): Promise<VolumeBaseInfo> {
	const resp = await requestApi(
		{ method: 'post', url: '/rest/v1/getVolumeById', data: payload },
		VolumeBaseInfoResponseSchema,
		{ errorTitle: 'Failed to load volume' }
	);
	return resp.item;
}

export function removeVolume(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeVolume', data: payload },
		{
			errorTitle: 'Failed to delete volume'
		}
	);
}
