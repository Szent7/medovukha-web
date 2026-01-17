import { z } from 'zod';

export const ContainerBaseInfoSchema = z.object({
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
export type ContainerBaseInfo = z.infer<typeof ContainerBaseInfoSchema>;

export const ListContainerBaseInfoSchema = z.object({
	items: z.array(ContainerBaseInfoSchema)
});
export type ListContainerBaseInfo = z.infer<typeof ListContainerBaseInfoSchema>;

export const ContainerIdMessageSchema = z.object({ id: z.string() });
export type ContainerIdMessage = z.infer<typeof ContainerIdMessageSchema>;

export const ActorSchema = z.object({ id: z.string() });
export const ContainerEventSchema = z.object({
	type: z.string(),
	action: z.string(),
	actor: ActorSchema
});
export type ContainerEvent = z.infer<typeof ContainerEventSchema>;
