import { z } from 'zod';

export const DeployFromGitSchema = z.object({
	url: z.string(),
	dockerfile: z.string(),
	docker_compose: z.string(),
	docker_run: z.string()
});
export type DeployFromGitInfo = z.infer<typeof DeployFromGitSchema>;

export const BuildIdSchema = z.object({ build_id: z.string() });
export type BuildId = z.infer<typeof BuildIdSchema>;
