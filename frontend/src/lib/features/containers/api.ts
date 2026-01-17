import { requestApi, requestApiNoData } from '$lib/core/http/request';
import {
	ListContainerBaseInfoSchema,
	type ContainerIdMessage,
	type ListContainerBaseInfo
} from './types';

export function getContainerList(): Promise<ListContainerBaseInfo> {
	return requestApi(
		{ method: 'get', url: '/rest/v1/getContainerList' },
		ListContainerBaseInfoSchema,
		{
			errorTitle: 'Не удалось загрузить список контейнеров'
		}
	);
}

export function startContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/startContainerById', data: payload },
		{
			errorTitle: 'Не удалось запустить контейнер'
		}
	);
}

export function stopContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/stopContainerById', data: payload },
		{
			errorTitle: 'Не удалось остановить контейнер'
		}
	);
}

export function killContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/killContainerById', data: payload },
		{
			errorTitle: 'Не удалось kill контейнер'
		}
	);
}

export function restartContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/restartContainerById', data: payload },
		{
			errorTitle: 'Не удалось перезапустить контейнер'
		}
	);
}

export function pauseContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/pauseContainerById', data: payload },
		{
			errorTitle: 'Не удалось поставить на паузу контейнер'
		}
	);
}

export function unpauseContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/unpauseContainerById', data: payload },
		{
			errorTitle: 'Не удалось снять с паузы контейнер'
		}
	);
}

export function removeContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeContainerById', data: payload },
		{
			errorTitle: 'Не удалось удалить контейнер'
		}
	);
}
