<script lang="ts" context="module">
	import { z } from 'zod';
	//
	//
	// Response
	//
	//
	export const createResponseScheme = <T extends z.ZodTypeAny>(dataScheme: T) =>
		z.object({
			success: z.boolean(),
			error: APIErrorSchema.optional(),
			data: dataScheme.optional()
		});

	export const APIErrorSchema = z.object({
		code: z.string(),
		message: z.string()
	});
	//
	//
	// ContainerIdMessage
	//
	//
	export const ContainerIdMessageScheme = z.object({
		id: z.string()
	});
	export type ContainerIdMessage = z.infer<typeof ContainerIdMessageScheme>;

	export const BuildIDScheme = z.object({
		build_id: z.string()
	});
	export type BuildIDMessage = z.infer<typeof BuildIDScheme>;
	//
	//
	// ContainerBaseInfo
	//
	//
	export const ContainerBaseInfoScheme = z.object({
		id: z.string(),
		names: z.array(z.string()),
		image_name: z.string(),
		ports: z
			.array(
				z
					.object({
						ip: z.string(),
						private_port: z.number(),
						public_port: z.number(),
						type: z.string()
					})
					.optional()
			)
			.optional(),
		created: z.number(),
		state: z.string()
	});
	export type ContainerBaseInfo = z.infer<typeof ContainerBaseInfoScheme>;

	export const ListContainerBaseInfoScheme = z.object({
		items: z.array(ContainerBaseInfoScheme)
	});
	export type ListContainerBaseInfo = z.infer<typeof ListContainerBaseInfoScheme>;
	//
	//
	// ImageBaseInfo
	//
	//
	export const ImageBaseInfoScheme = z.object({
		id: z.string(),
		tags: z.array(z.string()),
		size: z.number(),
		created: z.number()
	});
	export type ImageBaseInfo = z.infer<typeof ImageBaseInfoScheme>;

	export const ListImageBaseInfoScheme = z.object({
		items: z.array(ImageBaseInfoScheme)
	});
	export type ListImageBaseInfo = z.infer<typeof ListImageBaseInfoScheme>;
	//
	//
	// NetworkBaseInfo
	//
	//
	export const NetworkBaseInfoScheme = z.object({
		name: z.string(),
		id: z.string(),
		driver: z.string(),
		enable_ipv6: z.boolean().optional(),
		ipam_driver: z.string(),
		subnet: z.array(z.string()).optional(),
		gateway: z.array(z.string()).optional(),
		attachable: z.boolean().optional(),
		docker_network: z.boolean().optional()
	});
	export type NetworkBaseInfo = z.infer<typeof NetworkBaseInfoScheme>;

	export const ListNetworkBaseInfoScheme = z.object({
		items: z.array(NetworkBaseInfoScheme)
	});
	export type ListNetworkBaseInfo = z.infer<typeof ListNetworkBaseInfoScheme>;
	//
	//
	// VolumeBaseInfo
	//
	//
	export const VolumeBaseInfoScheme = z.object({
		name: z.string(),
		driver: z.string(),
		mountpoint: z.string(),
		created: z.string()
	});
	export type VolumeBaseInfo = z.infer<typeof VolumeBaseInfoScheme>;

	export const ListVolumeBaseInfoScheme = z.object({
		items: z.array(VolumeBaseInfoScheme)
	});
	export type ListVolumeBaseInfo = z.infer<typeof ListVolumeBaseInfoScheme>;
	//
	//
	// DeployFromGitInfo
	//
	//
	export const DeployFromGitScheme = z.object({
		url: z.string(),
		dockerfile: z.string(),
		docker_compose: z.string(),
		docker_run: z.string()
	});
	export type DeployFromGitInfo = z.infer<typeof DeployFromGitScheme>;
</script>
