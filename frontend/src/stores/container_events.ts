import { writable } from 'svelte/store';
import { string, z } from 'zod';

export const ActorSchema = z.object({
	id: z.string()
	//attributes: z.record(z.string(), z.string())
});
export type Actor = z.infer<typeof ActorSchema>;

export const containerEventSchema = z.object({
	type: z.string(),
	action: z.string(),
	actor: ActorSchema
});
export type ContainerEvent = z.infer<typeof containerEventSchema>;

export const containerEventsStore = writable<ContainerEvent | null>(null);
