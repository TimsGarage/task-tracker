<script>
	let { message, confirmLabel = 'Delete', onconfirm, oncancel } = $props();

	let working = $state(false);

	async function confirm() {
		if (working) {
			return;
		}

		working = true;

		try {
			await onconfirm();
		} catch {
			// The parent application displays the error.
		} finally {
			working = false;
		}
	}
</script>

<div class="flex min-w-0 flex-1 flex-wrap items-center justify-between gap-3">
	<span class="text-on-surface text-xs">
		{message}
	</span>

	<div class="flex items-center gap-2">
		<button
			type="button"
			class="bg-error text-on-error hover:bg-error/90 rounded-lg px-3 py-2 text-xs font-semibold transition-colors disabled:opacity-50"
			disabled={working}
			onclick={confirm}
		>
			{working ? 'Deleting…' : confirmLabel}
		</button>

		<button
			type="button"
			class="text-on-surface-variant hover:bg-surface-container-high rounded-lg px-3 py-2 text-xs transition-colors"
			disabled={working}
			onclick={oncancel}
		>
			Cancel
		</button>
	</div>
</div>
