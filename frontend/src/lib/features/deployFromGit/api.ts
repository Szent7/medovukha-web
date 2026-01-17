import { requestApi } from '$lib/core/http/request';
import { BuildIdSchema, type BuildId, type DeployFromGitInfo } from './types';

export function createFromGit(payload: DeployFromGitInfo): Promise<BuildId> {
	return requestApi(
		{ method: 'post', url: '/rest/v1/createFromGit', data: payload },
		BuildIdSchema,
		{
			errorTitle: 'Не удалось запустить сборку из Git'
		}
	);
}
