<script>
	let { project } = $props();
</script>

<aside
	class="bg-surface-container-lowest sticky top-24 hidden flex-col gap-6 rounded-xl p-6 shadow-sm xl:col-span-5 xl:flex"
>
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-2">
			<span class="bg-primary h-2.5 w-2.5 rounded-full"></span>

			<span class="text-secondary font-label text-xs font-semibold tracking-wider uppercase">
				{project?.title ?? 'Task details'}
			</span>
		</div>

		<div class="flex items-center gap-1.5">
			<button
				class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg transition-colors"
				type="button"
				title="Favorite"
			>
				<span class="material-symbols-outlined text-lg"> star </span>
			</button>

			<button
				class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg transition-colors"
				type="button"
				title="Options"
			>
				<span class="material-symbols-outlined text-lg"> more_horiz </span>
			</button>
		</div>
	</div>

	<div class="flex items-start gap-3">
		<div class="border-outline mt-1 h-6 w-6 shrink-0 rounded-full border-2"></div>

		<div>
			<h2 class="font-headline text-on-surface text-lg leading-snug font-semibold">
				Select a task to inspect
			</h2>

			<p class="text-on-surface-variant mt-1 text-xs">
				Detailed task attributes will appear here in a future version.
			</p>
		</div>
	</div>

	<div class="bg-surface-container-low grid grid-cols-2 gap-3 rounded-xl p-3.5">
		<div class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Due date
			</span>

			<div
				class="bg-surface-container-lowest text-outline flex items-center gap-2 rounded-lg px-3 py-2 text-xs"
			>
				<span class="material-symbols-outlined text-sm"> calendar_today </span>

				Not available
			</div>
		</div>

		<div class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Priority
			</span>

			<div
				class="bg-surface-container-lowest text-outline flex items-center gap-2 rounded-lg px-3 py-2 text-xs"
			>
				<span class="material-symbols-outlined text-sm"> flag </span>

				Not available
			</div>
		</div>

		<div class="col-span-2 flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Project
			</span>

			<div
				class="bg-surface-container-lowest text-on-surface flex items-center gap-2 rounded-lg px-3 py-2 text-xs"
			>
				<span class="bg-primary h-2 w-2 rounded-full"></span>

				{project?.title ?? 'No project selected'}
			</div>
		</div>
	</div>

	<section class="space-y-2">
		<span class="text-secondary text-xs font-semibold tracking-wider uppercase"> Field notes </span>

		<div
			class="bg-surface-container-low text-on-surface-variant rounded-xl p-4 text-xs leading-relaxed"
		>
			Notes, attachments, subtasks, activity, and scheduling can be implemented here later.
		</div>
	</section>

	<div class="border-outline-variant border-t pt-4">
		<button
			type="button"
			class="bg-primary text-on-primary w-full rounded-xl px-4 py-2.5 text-xs font-medium opacity-50"
			disabled
		>
			Save changes
		</button>
	</div>
</aside>
