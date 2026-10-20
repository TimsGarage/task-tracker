<script>
	import ProgressCard from './ProgressCard.svelte';

	let { project, completed = 0, total = 0, percentage = 0 } = $props();

	const formattedDate = new Intl.DateTimeFormat('en', {
		weekday: 'long',
		month: 'long',
		day: 'numeric'
	}).format(new Date());
</script>

<div class="mb-8 flex flex-col justify-between gap-6 lg:flex-row lg:items-end">
	<div>
		<div class="mb-2 flex flex-wrap items-center gap-2.5">
			<span
				class="bg-primary-fixed text-on-primary-fixed inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-semibold tracking-wide"
			>
				<span class="material-symbols-outlined text-sm"> filter_vintage </span>

				Project Rhythm
			</span>

			<span class="text-secondary font-label text-xs font-medium tracking-wider uppercase">
				{formattedDate}
			</span>
		</div>

		<h1 class="font-headline text-on-surface text-3xl font-semibold tracking-tight sm:text-4xl">
			{project?.title ?? 'Your Tasks'}
		</h1>

		<p class="text-on-surface-variant mt-1.5 text-sm">
			Ground yourself in deliberate, thoughtful progress.
		</p>
	</div>

	<ProgressCard {completed} {total} {percentage} />
</div>
