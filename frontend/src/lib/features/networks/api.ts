import { requestApi, requestApiNoData } from '$lib/core/http/request';
import {
	ListNetworkBaseInfoSchema,
	NetworkBaseInfoResponseSchema,
	type NetworkBaseInfo,
	type ListNetworkBaseInfo,
	type RemoveIdMessage
} from './types';

export function getNetworkList(): Promise<ListNetworkBaseInfo> {
	return requestApi({ method: 'get', url: '/rest/v1/getNetworkList' }, ListNetworkBaseInfoSchema, {
		errorTitle: 'Failed to load network list'
	});
}

export async function getNetworkById(payload: { id: string }): Promise<NetworkBaseInfo> {
	const resp = await requestApi(
		{ method: 'post', url: '/rest/v1/getNetworkById', data: payload },
		NetworkBaseInfoResponseSchema,
		{ errorTitle: 'Failed to load network' }
	);
	return resp.item;
}

export function removeNetwork(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeNetwork', data: payload },
		{
			errorTitle: 'Failed to delete network'
		}
	);
}
