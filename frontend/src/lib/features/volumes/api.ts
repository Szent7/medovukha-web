import { requestApi, requestApiNoData } from '$lib/core/http/request';
import { ListVolumeBaseInfoSchema, type ListVolumeBaseInfo, type RemoveIdMessage } from './types';

export function getVolumeList(): Promise<ListVolumeBaseInfo> {
	return requestApi({ method: 'get', url: '/rest/v1/getVolumeList' }, ListVolumeBaseInfoSchema, {
		errorTitle: 'Не удалось загрузить список томов'
	});
}

export function removeVolume(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeVolume', data: payload },
		{
			errorTitle: 'Не удалось удалить том'
		}
	);
}
