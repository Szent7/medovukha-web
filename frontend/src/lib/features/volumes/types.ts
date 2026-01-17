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

export const RemoveIdMessageSchema = z.object({
	id: z.string(),
	force: z.boolean().optional()
});
export type RemoveIdMessage = z.infer<typeof RemoveIdMessageSchema>;
