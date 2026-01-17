import { requestApi, requestApiNoData } from '$lib/core/http/request';
import { ListImageBaseInfoSchema, type ListImageBaseInfo, type RemoveIdMessage } from './types';

export function getImageList(): Promise<ListImageBaseInfo> {
	return requestApi({ method: 'get', url: '/rest/v1/getImageList' }, ListImageBaseInfoSchema, {
		errorTitle: 'Не удалось загрузить список образов'
	});
}

export function removeImage(payload: RemoveIdMessage): Promise<void> {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeImage', data: payload },
		{
			errorTitle: 'Не удалось удалить образ'
		}
	);
}
