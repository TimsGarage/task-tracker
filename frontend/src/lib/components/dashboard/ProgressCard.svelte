<script>
	let { completed = 0, total = 0, percentage = 0 } = $props();

	const circumference = 125.6;

	let dashOffset = $derived(circumference - (circumference * percentage) / 100);
</script>

<div class="bg-surface-container flex items-center gap-5 rounded-xl p-4 shadow-sm">
	<div class="relative flex h-14 w-14 shrink-0 items-center justify-center">
		<svg class="h-14 w-14 -rotate-45 transform" viewBox="0 0 48 48" aria-hidden="true">
			<circle
				class="text-surface-container-high"
				cx="24"
				cy="24"
				fill="transparent"
				r="20"
				stroke="currentColor"
				stroke-width="4"
			></circle>

			<circle
				class="text-primary transition-all duration-700"
				cx="24"
				cy="24"
				fill="transparent"
				r="20"
				stroke="currentColor"
				stroke-dasharray={circumference}
				stroke-dashoffset={dashOffset}
				stroke-linecap="round"
				stroke-width="4.5"
			></circle>
		</svg>

		<span class="text-on-surface absolute text-xs font-semibold">
			{percentage}%
		</span>
	</div>

	<div>
		<div class="flex items-center gap-2">
			<span class="font-headline text-on-surface text-sm font-semibold">
				{completed} of {total} completed
			</span>

			<span class="material-symbols-outlined text-primary text-base"> spa </span>
		</div>

		<p class="text-on-surface-variant mt-0.5 text-xs">
			{#if total === 0}
				Add your first task and start cultivating progress.
			{:else if percentage === 100}
				Everything in this project has been completed.
			{:else if percentage >= 50}
				Steady progress. You are past the halfway point.
			{:else}
				A deliberate start is still meaningful progress.
			{/if}
		</p>
	</div>
</div>
