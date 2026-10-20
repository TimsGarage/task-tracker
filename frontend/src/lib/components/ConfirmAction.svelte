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
			// The parent displays the API error.
		} finally {
			working = false;
		}
	}
</script>

<div class="confirmation">
	<span>{message}</span>

	<div class="actions">
		<button type="button" class="danger" onclick={confirm} disabled={working}>
			{working ? 'Deleting…' : confirmLabel}
		</button>

		<button type="button" onclick={oncancel} disabled={working}> Cancel </button>
	</div>
</div>

<style>
	.confirmation {
		display: flex;
		flex: 1;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		width: 100%;
	}

	.actions {
		display: flex;
		flex-shrink: 0;
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

	button.danger {
		color: white;
		border-color: #dc2626;
		background: #dc2626;
	}

	button:disabled {
		cursor: not-allowed;
		opacity: 0.5;
	}

	@media (max-width: 600px) {
		.confirmation {
			align-items: stretch;
			flex-direction: column;
		}
	}
</style>
