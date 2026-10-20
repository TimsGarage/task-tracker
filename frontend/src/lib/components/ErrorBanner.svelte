<script>
	let { message, ondismiss } = $props();
</script>

{#if message}
	<div class="error" role="alert">
		<span>{message}</span>

		<button type="button" onclick={ondismiss} aria-label="Dismiss error"> × </button>
	</div>
{/if}

<style>
	.error {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 1rem;
		margin-bottom: 1rem;
		padding: 0.75rem 1rem;
		color: #991b1b;
		border: 1px solid #fecaca;
		border-radius: 8px;
		background: #fee2e2;
	}

	button {
		padding: 0.2rem 0.5rem;
		color: inherit;
		border: 0;
		background: transparent;
		font-size: 1.2rem;
		cursor: pointer;
	}
</style>
