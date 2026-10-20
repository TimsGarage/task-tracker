<script>
	import ConfirmAction from '../ui/ConfirmAction.svelte';
	import InlineEditor from '../ui/InlineEditor.svelte';

	let { todo, projectTitle, selected = false, onselect, ontoggle, onrename, ondelete } = $props();

	let mode = $state('view');
	let toggling = $state(false);

	async function toggle() {
		if (toggling) return;

		toggling = true;

		try {
			await ontoggle(todo.id, !todo.done);
		} catch {
			// Handled by the app-level banner.
		} finally {
			toggling = false;
		}
	}

	async function saveRename(title) {
		if (title !== todo.title) {
			await onrename(todo.id, title);
		}

		mode = 'view';
	}

	async function confirmDelete() {
		await ondelete(todo.id);
		mode = 'view';
	}
</script>

<article
	class={[
		'group relative rounded-xl p-4 shadow-sm transition-all',
		selected
			? 'bg-surface-container-lowest ring-primary/30 ring-2'
			: todo.done
				? 'bg-surface-container-low opacity-75 hover:opacity-100'
				: 'bg-surface-container-low hover:bg-surface-container-lowest'
	]}
>
	{#if selected}
		<div class="bg-primary absolute top-1/4 bottom-0 -left-1 h-1/2 w-1.5 rounded-xl"></div>
	{/if}

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
			onconfirm={confirmDelete}
			oncancel={() => (mode = 'view')}
		/>
	{:else}
		<div class="flex min-w-0 items-start justify-between gap-4">
			<div class="flex min-w-0 items-start gap-3.5">
				<button
					type="button"
					class={[
						'mt-0.5 flex h-5 w-5 shrink-0 items-center justify-center rounded-full transition-all active:scale-90',
						todo.done
							? 'bg-primary text-on-primary'
							: 'border-outline hover:border-primary border-2'
					]}
					disabled={toggling}
					aria-label={todo.done ? 'Mark as incomplete' : 'Mark as completed'}
					onclick={toggle}
				>
					{#if todo.done}
						<span class="material-symbols-outlined text-sm font-bold"> check </span>
					{/if}
				</button>

				<button type="button" class="min-w-0 text-left" onclick={() => onselect(todo.id)}>
					<span
						class={[
							'block max-w-full truncate text-sm',
							todo.done ? 'text-outline font-medium line-through' : 'text-on-surface font-semibold'
						]}
					>
						{todo.title}
					</span>

					<span class="mt-2 flex flex-wrap items-center gap-2">
						<span
							class="bg-primary-fixed text-on-primary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
						>
							{projectTitle}
						</span>

						{#if todo.dueDate}
							<span
								class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px]"
							>
								<span class="material-symbols-outlined text-xs"> schedule </span>

								{new Intl.DateTimeFormat('en', {
									month: 'short',
									day: 'numeric',
									hour: 'numeric',
									minute: '2-digit'
								}).format(new Date(todo.dueDate))}
							</span>
						{/if}

						{#if todo.subtasks?.length}
							<span
								class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px]"
							>
								<span class="material-symbols-outlined text-xs"> checklist </span>

								{todo.subtasks.length}
							</span>
						{/if}
					</span>
				</button>
			</div>

			<div
				class="flex shrink-0 items-center opacity-0 transition-opacity group-focus-within:opacity-100 group-hover:opacity-100"
			>
				<button
					type="button"
					class="text-on-surface-variant hover:bg-surface-container-high flex h-8 w-8 items-center justify-center rounded-lg"
					title="Rename task"
					onclick={() => (mode = 'rename')}
				>
					<span class="material-symbols-outlined text-base"> edit </span>
				</button>

				<button
					type="button"
					class="text-error hover:bg-error-container flex h-8 w-8 items-center justify-center rounded-lg"
					title="Delete task"
					onclick={() => (mode = 'delete')}
				>
					<span class="material-symbols-outlined text-base"> delete </span>
				</button>
			</div>
		</div>
	{/if}
</article>
