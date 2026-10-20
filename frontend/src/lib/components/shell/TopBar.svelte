<header
	class="bg-surface/80 fixed top-0 right-0 left-0 z-40 flex h-16 items-center justify-between gap-4 px-4 shadow-[0_1px_8px_rgba(46,50,48,0.04)] backdrop-blur-xl lg:left-72 lg:px-8"
>
	<div class="flex w-full max-w-96 items-center">
		<div class="relative flex w-full items-center">
			<span
				class="material-symbols-outlined text-on-surface-variant pointer-events-none absolute left-3 text-lg"
			>
				search
			</span>

			<input
				class="bg-surface-container-low text-on-surface placeholder:text-on-surface-variant focus:ring-primary/40 w-full rounded-xl border-0 py-2 pr-4 pl-9 text-sm transition-all focus:ring-2 focus:outline-none"
				placeholder="Search tasks, projects, tags..."
				aria-label="Search"
				type="search"
			/>
		</div>
	</div>

	<div class="flex shrink-0 items-center gap-2 sm:gap-4">
		<button
			class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface hidden h-9 w-9 items-center justify-center rounded-xl transition-colors sm:flex"
			type="button"
			aria-label="Notifications"
		>
			<span class="material-symbols-outlined text-xl"> notifications </span>
		</button>

		<button
			class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface hidden h-9 w-9 items-center justify-center rounded-xl transition-colors sm:flex"
			type="button"
			aria-label="Filters"
		>
			<span class="material-symbols-outlined text-xl"> tune </span>
		</button>

		<div class="flex items-center gap-3 pl-1">
			<div
				class="bg-primary-fixed text-primary flex h-8 w-8 items-center justify-center rounded-full"
			>
				<span class="material-symbols-outlined text-lg"> person </span>
			</div>

			<div class="hidden flex-col text-left xl:flex">
				<span class="text-on-surface text-xs leading-tight font-semibold"> Local Gardener </span>

				<span class="text-on-surface-variant text-[11px] leading-tight"> Personal workspace </span>
			</div>
		</div>
	</div>
</header>
