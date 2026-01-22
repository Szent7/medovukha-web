import { z } from 'zod';

export const NetworkBaseInfoSchema = z.object({
	name: z.string(),
	id: z.string(),
	driver: z.string(),
	enable_ipv6: z.boolean().optional(),
	ipam_driver: z.string(),
	subnet: z.array(z.string()).optional(),
	gateway: z.array(z.string()).optional(),
	attachable: z.boolean().optional(),
	docker_network: z.boolean().optional(),
	is_used: z.boolean().optional()
});
export type NetworkBaseInfo = z.infer<typeof NetworkBaseInfoSchema>;

export const ListNetworkBaseInfoSchema = z.object({
	items: z.array(NetworkBaseInfoSchema)
});
export type ListNetworkBaseInfo = z.infer<typeof ListNetworkBaseInfoSchema>;

export const NetworkBaseInfoResponseSchema = z.object({
	item: NetworkBaseInfoSchema
});
export type NetworkBaseInfoResponse = z.infer<typeof NetworkBaseInfoResponseSchema>;

export const NetworkEventSchema = z.object({
	type: z.string(),
	action: z.string(),
	actor: z.object({
		id: z.string()
	})
});
export type NetworkEvent = z.infer<typeof NetworkEventSchema>;

export const RemoveIdMessageSchema = z.object({ id: z.string() });
export type RemoveIdMessage = z.infer<typeof RemoveIdMessageSchema>;
