<script>
	let { message, ondismiss } = $props();
</script>

{#if message}
	<div
		class="bg-error-container text-on-error-container mb-6 flex items-center justify-between gap-4 rounded-xl px-4 py-3 text-sm shadow-sm"
		role="alert"
	>
		<div class="flex items-center gap-2">
			<span class="material-symbols-outlined text-lg"> error </span>

			<span>{message}</span>
		</div>

		<button
			type="button"
			class="hover:bg-error/10 flex h-8 w-8 shrink-0 items-center justify-center rounded-lg transition-colors"
			onclick={ondismiss}
			aria-label="Dismiss error"
		>
			<span class="material-symbols-outlined text-lg"> close </span>
		</button>
	</div>
{/if}
