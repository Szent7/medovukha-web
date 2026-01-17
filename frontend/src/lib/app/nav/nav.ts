export type NavItem = {
	href: string;
	label: string;
};

export const NAV: NavItem[] = [
	{ href: '/', label: 'Home' },
	{ href: '/containers', label: 'Containers' },
	{ href: '/images', label: 'Images' },
	{ href: '/networks', label: 'Networks' },
	{ href: '/volumes', label: 'Volumes' }
];
