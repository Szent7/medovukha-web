<script lang="ts" context="module">
	import { z } from 'zod';
	//
	//
	// ContainerIdMessage
	//
	//
	export const ContainerIdMessageScheme = z.object({
		id: z.string()
	});
	export type ContainerIdMessage = z.infer<typeof ContainerIdMessageScheme>;
	//
	//
	// ContainerBaseInfo
	//
	//
	export const ContainerBaseInfoScheme = z.object({
		id: z.string(),
		names: z.array(z.string()),
		imageName: z.string(),
		ports: z
			.array(
				z
					.object({
						ip: z.string(),
						privatePort: z.number(),
						publicPort: z.number(),
						type: z.string()
					})
					.optional()
			)
			.optional(),
		created: z.string(),
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
		size: z.string(),
		created: z.string()
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
		enableIPv6: z.boolean().optional(),
		ipamDriver: z.string(),
		subnet: z.array(z.string()).optional(),
		gateway: z.array(z.string()).optional(),
		attachable: z.boolean().optional(),
		dockerNetwork: z.boolean().optional()
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
</script>
