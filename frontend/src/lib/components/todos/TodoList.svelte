<script>
	import TodoItem from './TodoItem.svelte';

	let { todos, projectTitle, focusToken = 0, oncreate, ontoggle, onrename, ondelete } = $props();

	let newTodoTitle = $state('');
	let creating = $state(false);
	let inputElement = $state();
	let previousFocusToken = $state(focusToken);

	let openTodos = $derived(todos.filter((todo) => !todo.done));

	let completedTodos = $derived(todos.filter((todo) => todo.done));

	let canCreate = $derived(newTodoTitle.trim().length > 0 && !creating);

	$effect(() => {
		if (focusToken !== previousFocusToken) {
			previousFocusToken = focusToken;
			inputElement?.focus();
		}
	});

	async function submit(event) {
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
			// The app-level error banner handles the error.
		} finally {
			creating = false;
		}
	}
</script>

<div class="mb-8">
	<form
		class="bg-surface-container-lowest relative flex items-center gap-3 rounded-xl p-2 shadow-sm transition-all focus-within:shadow-md"
		onsubmit={submit}
	>
		<div
			class="bg-primary-fixed text-primary ml-1 flex h-9 w-9 shrink-0 items-center justify-center rounded-full"
		>
			<span class="material-symbols-outlined text-xl"> add </span>
		</div>

		<input
			bind:this={inputElement}
			bind:value={newTodoTitle}
			class="text-on-surface placeholder:text-outline w-full bg-transparent py-1.5 text-sm focus:outline-none"
			placeholder="Add a mindful task... Press Enter to cultivate"
			aria-label="New todo title"
			disabled={creating}
		/>

		<button
			class="bg-primary text-on-primary hover:bg-primary/90 rounded-lg px-3.5 py-2 text-xs font-medium transition-colors disabled:opacity-50"
			type="submit"
			disabled={!canCreate}
		>
			{creating ? 'Cultivating…' : 'Cultivate'}
		</button>
	</form>
</div>

<div class="flex flex-col gap-8">
	<section class="space-y-3">
		<div class="flex items-center justify-between pb-1">
			<div class="flex items-center gap-2">
				<span class="material-symbols-outlined text-primary text-lg"> wb_sunny </span>

				<h2 class="font-headline text-on-surface text-base font-semibold tracking-tight">
					Active Tasks
				</h2>
			</div>

			<span
				class="bg-surface-container text-on-surface-variant rounded-full px-2 py-0.5 text-xs font-semibold"
			>
				{openTodos.length} remaining
			</span>
		</div>

		<div class="space-y-2.5">
			{#each openTodos as todo (todo.id)}
				<TodoItem {todo} {projectTitle} {ontoggle} {onrename} {ondelete} />
			{/each}

			{#if openTodos.length === 0}
				<div
					class="bg-surface-container-low text-on-surface-variant rounded-xl px-5 py-8 text-center text-sm"
				>
					No active tasks. The garden is calm.
				</div>
			{/if}
		</div>
	</section>

	{#if completedTodos.length > 0}
		<section class="space-y-3">
			<div class="flex items-center justify-between pb-1">
				<div class="flex items-center gap-2">
					<span class="material-symbols-outlined text-tertiary text-lg"> check_circle </span>

					<h2 class="font-headline text-on-surface text-base font-semibold tracking-tight">
						Completed
					</h2>
				</div>

				<span
					class="bg-surface-container text-secondary rounded-full px-2 py-0.5 text-xs font-semibold"
				>
					{completedTodos.length} done
				</span>
			</div>

			<div class="space-y-2">
				{#each completedTodos as todo (todo.id)}
					<TodoItem {todo} {projectTitle} {ontoggle} {onrename} {ondelete} />
				{/each}
			</div>
		</section>
	{/if}
</div>
