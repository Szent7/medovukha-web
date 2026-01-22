import { z } from 'zod';

export const ImageBaseInfoSchema = z.object({
	id: z.string(),
	tags: z.array(z.string()),
	size: z.number(),
	created: z.number(),
	is_used: z.boolean().optional()
});
export type ImageBaseInfo = z.infer<typeof ImageBaseInfoSchema>;

export const ListImageBaseInfoSchema = z.object({
	items: z.array(ImageBaseInfoSchema)
});
export type ListImageBaseInfo = z.infer<typeof ListImageBaseInfoSchema>;

export const ImageBaseInfoResponseSchema = z.object({
	item: ImageBaseInfoSchema
});
export type ImageBaseInfoResponse = z.infer<typeof ImageBaseInfoResponseSchema>;

export const ImageEventSchema = z.object({
	type: z.string(),
	action: z.string(),
	actor: z.object({
		id: z.string()
	})
});
export type ImageEvent = z.infer<typeof ImageEventSchema>;

export const RemoveIdMessageSchema = z.object({
	id: z.string(),
	force: z.boolean().optional()
});
export type RemoveIdMessage = z.infer<typeof RemoveIdMessageSchema>;
