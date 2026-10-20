<script>
	import ConfirmAction from './ConfirmAction.svelte';
	import InlineEditor from './InlineEditor.svelte';

	let { todo, ontoggle, onrename, ondelete } = $props();

	let mode = $state('view');
	let toggling = $state(false);

	async function toggle() {
		if (toggling) {
			return;
		}

		toggling = true;

		try {
			await ontoggle(todo.id, !todo.done);
		} catch {
			// The application displays the error.
		} finally {
			toggling = false;
		}
	}

	async function saveRename(title) {
		if (title === todo.title) {
			mode = 'view';
			return;
		}

		await onrename(todo.id, title);
		mode = 'view';
	}

	async function confirmDelete() {
		await ondelete(todo.id);
		mode = 'view';
	}
</script>

<li class:done={todo.done}>
	{#if mode === 'rename'}
		<InlineEditor
			initialValue={todo.title}
			label="Todo title"
			onsave={saveRename}
			oncancel={() => (mode = 'view')}
		/>
	{:else if mode === 'delete'}
		<ConfirmAction
			message={`Delete "${todo.title}"?`}
			confirmLabel="Delete"
			onconfirm={confirmDelete}
			oncancel={() => (mode = 'view')}
		/>
	{:else}
		<button
			type="button"
			class="checkbox"
			class:checked={todo.done}
			onclick={toggle}
			disabled={toggling}
			aria-label={todo.done ? 'Mark as incomplete' : 'Mark as done'}
		>
			{todo.done ? '✓' : ''}
		</button>

		<button
			type="button"
			class="todo-title"
			onclick={() => (mode = 'rename')}
			title="Click to rename"
		>
			{todo.title}
		</button>

		<button type="button" class="secondary" onclick={() => (mode = 'rename')}> Rename </button>

		<button type="button" class="secondary danger" onclick={() => (mode = 'delete')}>
			Delete
		</button>
	{/if}
</li>

<style>
	li {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		min-height: 58px;
		padding: 0.6rem;
		border-radius: 10px;
		background: white;
	}

	button {
		font: inherit;
		cursor: pointer;
	}

	.checkbox {
		display: grid;
		flex: 0 0 28px;
		width: 28px;
		height: 28px;
		padding: 0;
		place-items: center;
		color: white;
		border: 2px solid #9ca3af;
		border-radius: 50%;
		background: white;
	}

	.checkbox.checked {
		border-color: #16a34a;
		background: #16a34a;
	}

	.checkbox:disabled {
		cursor: wait;
		opacity: 0.5;
	}

	.todo-title {
		flex: 1;
		min-width: 0;
		padding: 0.65rem;
		overflow: hidden;
		color: inherit;
		border: 0;
		background: transparent;
		text-align: left;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.done .todo-title {
		color: #9ca3af;
		text-decoration: line-through;
	}

	.secondary {
		padding: 0.6rem 0.75rem;
		border: 1px solid #d1d5db;
		border-radius: 7px;
		background: white;
	}

	.danger {
		color: #dc2626;
	}

	@media (max-width: 650px) {
		li {
			flex-wrap: wrap;
		}

		.todo-title {
			width: calc(100% - 40px);
		}
	}
</style>
