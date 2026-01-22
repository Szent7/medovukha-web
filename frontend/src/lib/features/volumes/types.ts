import { z } from 'zod';

export const VolumeBaseInfoSchema = z.object({
	name: z.string(),
	driver: z.string(),
	mountpoint: z.string(),
	created: z.string(),
	is_used: z.boolean().optional()
});
export type VolumeBaseInfo = z.infer<typeof VolumeBaseInfoSchema>;

export const ListVolumeBaseInfoSchema = z.object({
	items: z.array(VolumeBaseInfoSchema)
});
export type ListVolumeBaseInfo = z.infer<typeof ListVolumeBaseInfoSchema>;

export const VolumeBaseInfoResponseSchema = z.object({
	item: VolumeBaseInfoSchema
});
export type VolumeBaseInfoResponse = z.infer<typeof VolumeBaseInfoResponseSchema>;

export const VolumeEventSchema = z.object({
	type: z.string(),
	action: z.string(),
	actor: z.object({
		id: z.string()
	})
});
export type VolumeEvent = z.infer<typeof VolumeEventSchema>;

export const RemoveIdMessageSchema = z.object({
	id: z.string(),
	force: z.boolean().optional()
});
export type RemoveIdMessage = z.infer<typeof RemoveIdMessageSchema>;
