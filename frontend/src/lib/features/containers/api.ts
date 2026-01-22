import { requestApi, requestApiNoData } from '$lib/core/http/request';
import {
	ContainerBaseInfoResponseSchema,
	ListContainerBaseInfoSchema,
	type ContainerIdMessage,
	type ContainerBaseInfo,
	type ListContainerBaseInfo
} from './types';

export function getContainerList(): Promise<ListContainerBaseInfo> {
	return requestApi(
		{ method: 'get', url: '/rest/v1/getContainerList' },
		ListContainerBaseInfoSchema,
		{
			errorTitle: 'Failed to load container list'
		}
	);
}

export async function getContainerById(payload: ContainerIdMessage): Promise<ContainerBaseInfo> {
	const resp = await requestApi(
		{ method: 'post', url: '/rest/v1/getContainerById', data: payload },
		ContainerBaseInfoResponseSchema,
		{ errorTitle: 'Failed to load container' }
	);
	return resp.item;
}

export function startContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/startContainerById', data: payload },
		{
			errorTitle: 'Failed to start container'
		}
	);
}

export function stopContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/stopContainerById', data: payload },
		{
			errorTitle: 'Failed to stop container'
		}
	);
}

export function killContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/killContainerById', data: payload },
		{
			errorTitle: 'Failed to kill container'
		}
	);
}

export function restartContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/restartContainerById', data: payload },
		{
			errorTitle: 'Failed to restart container'
		}
	);
}

export function pauseContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/pauseContainerById', data: payload },
		{
			errorTitle: 'Failed to pause container'
		}
	);
}

export function unpauseContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/unpauseContainerById', data: payload },
		{
			errorTitle: 'Failed to unpause container'
		}
	);
}

export function removeContainerById(payload: ContainerIdMessage) {
	return requestApiNoData(
		{ method: 'post', url: '/rest/v1/removeContainerById', data: payload },
		{
			errorTitle: 'Failed to delete container'
		}
	);
}
