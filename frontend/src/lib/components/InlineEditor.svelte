<script>
	let { initialValue, label = 'Name', onsave, oncancel } = $props();

	let draft = $state(initialValue);
	let saving = $state(false);
	let inputElement = $state();

	let isValid = $derived(draft.trim().length > 0);

	$effect(() => {
		if (inputElement) {
			inputElement.focus();
			inputElement.select();
		}
	});

	async function submit(event) {
		event?.preventDefault();

		const value = draft.trim();

		if (!value || saving) {
			return;
		}

		saving = true;

		try {
			await onsave(value);
		} catch {
			// The parent displays the API error.
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

<form onsubmit={submit}>
	<input
		bind:this={inputElement}
		bind:value={draft}
		onkeydown={handleKeydown}
		aria-label={label}
		disabled={saving}
	/>

	<div class="actions">
		<button type="submit" class="primary" disabled={!isValid || saving}>
			{saving ? 'Saving…' : 'Save'}
		</button>

		<button type="button" onclick={oncancel} disabled={saving}> Cancel </button>
	</div>
</form>

<style>
	form {
		display: flex;
		flex: 1;
		gap: 0.4rem;
		width: 100%;
	}

	input {
		flex: 1;
		min-width: 0;
		padding: 0.6rem;
		border: 1px solid #9ca3af;
		border-radius: 7px;
		font: inherit;
	}

	.actions {
		display: flex;
		gap: 0.4rem;
	}

	button {
		padding: 0.6rem 0.75rem;
		border: 1px solid #d1d5db;
		border-radius: 7px;
		background: white;
		font: inherit;
		cursor: pointer;
	}

	button.primary {
		color: white;
		border-color: #2563eb;
		background: #2563eb;
	}

	button:disabled {
		cursor: not-allowed;
		opacity: 0.5;
	}

	@media (max-width: 600px) {
		form {
			flex-wrap: wrap;
		}

		input {
			flex-basis: 100%;
		}
	}
</style>
