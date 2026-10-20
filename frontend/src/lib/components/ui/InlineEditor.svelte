<script>
	let { initialValue, label = 'Name', onsave, oncancel } = $props();

	let draft = $state(initialValue);
	let saving = $state(false);
	let inputElement = $state();

	let canSave = $derived(draft.trim().length > 0 && !saving);

	$effect(() => {
		if (inputElement) {
			inputElement.focus();
			inputElement.select();
		}
	});

	async function submit(event) {
		event.preventDefault();

		const value = draft.trim();

		if (!value || saving) {
			return;
		}

		saving = true;

		try {
			await onsave(value);
		} catch {
			// The parent application displays the error.
		} finally {
			saving = false;
		}
	}

	function handleKeydown(event) {
		if (event.key === 'Escape') {
			event.preventDefault();
			oncancel();
		}
	}
</script>

<form class="flex min-w-0 flex-1 flex-wrap items-center gap-2" onsubmit={submit}>
	<input
		bind:this={inputElement}
		bind:value={draft}
		class="bg-surface-container-lowest text-on-surface focus:ring-primary/30 min-w-0 flex-1 rounded-lg border-0 px-3 py-2 text-sm focus:ring-2 focus:outline-none"
		aria-label={label}
		disabled={saving}
		onkeydown={handleKeydown}
	/>

	<button
		type="submit"
		class="bg-primary text-on-primary hover:bg-primary/90 rounded-lg px-3 py-2 text-xs font-semibold transition-colors disabled:opacity-50"
		disabled={!canSave}
	>
		{saving ? 'Saving…' : 'Save'}
	</button>

	<button
		type="button"
		class="text-on-surface-variant hover:bg-surface-container-high rounded-lg px-3 py-2 text-xs transition-colors"
		disabled={saving}
		onclick={oncancel}
	>
		Cancel
	</button>
</form>
