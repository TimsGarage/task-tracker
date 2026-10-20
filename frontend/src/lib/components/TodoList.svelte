<script>
	import TodoItem from './TodoItem.svelte';

	let { todos, oncreate, ontoggle, onrename, ondelete } = $props();

	let newTodoTitle = $state('');
	let creating = $state(false);

	let canCreate = $derived(newTodoTitle.trim().length > 0 && !creating);

	async function create(event) {
		event.preventDefault();

		const title = newTodoTitle.trim();

		if (!title || creating) {
			return;
		}

		creating = true;

		try {
			await oncreate(title);
			newTodoTitle = '';
		} catch {
			// The application displays the error.
		} finally {
			creating = false;
		}
	}
</script>

<form class="new-todo" onsubmit={create}>
	<input
		bind:value={newTodoTitle}
		placeholder="What needs doing?"
		aria-label="New todo title"
		disabled={creating}
	/>

	<button type="submit" disabled={!canCreate}>
		{creating ? 'Adding…' : 'Add todo'}
	</button>
</form>

<ul>
	{#each todos as todo (todo.id)}
		<TodoItem {todo} {ontoggle} {onrename} {ondelete} />
	{/each}
</ul>

{#if todos.length === 0}
	<p class="empty">This project has no todos yet.</p>
{/if}

<style>
	.new-todo {
		display: flex;
		gap: 0.75rem;
	}

	input {
		flex: 1;
		min-width: 0;
		padding: 0.8rem;
		border: 1px solid #d1d5db;
		border-radius: 8px;
		font: inherit;
	}

	.new-todo button {
		padding: 0.8rem 1rem;
		color: white;
		border: 1px solid #2563eb;
		border-radius: 8px;
		background: #2563eb;
		font: inherit;
		cursor: pointer;
	}

	button:disabled {
		cursor: not-allowed;
		opacity: 0.5;
	}

	ul {
		display: grid;
		gap: 0.6rem;
		padding: 0;
		list-style: none;
	}

	.empty {
		color: #6b7280;
	}

	@media (max-width: 500px) {
		.new-todo {
			align-items: stretch;
			flex-direction: column;
		}
	}
</style>
