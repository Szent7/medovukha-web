import { requestApi, requestApiNoData } from '$lib/core/http/request';
import { ListNetworkBaseInfoSchema, type ListNetworkBaseInfo, type RemoveIdMessage } from './types';

export function getNetworkList(): Promise<ListNetworkBaseInfo> {
	return requestApi({ method: 'get', url: '/rest/v1/getNetworkList' }, ListNetworkBaseInfoSchema, {
		errorTitle: 'Не удалось загрузить список сетей'
	});
}

export function removeNetwork(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeNetwork', data: payload },
		{
			errorTitle: 'Не удалось удалить сеть'
		}
	);
}
