<script>
	import ConfirmAction from '../ui/ConfirmAction.svelte';
	import InlineEditor from '../ui/InlineEditor.svelte';

	let { todo, projectTitle, ontoggle, onrename, ondelete } = $props();

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
			// The app-level error banner handles the error.
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

<article
	class={[
		'group rounded-xl p-4 shadow-sm transition-all',
		todo.done
			? 'bg-surface-container-low opacity-75 hover:opacity-100'
			: 'bg-surface-container-low hover:bg-surface-container-lowest'
	]}
>
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
						'mt-0.5 flex h-5 w-5 shrink-0 items-center justify-center rounded-full transition-all active:scale-90 disabled:opacity-50',
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

				<div class="min-w-0">
					<button
						type="button"
						class={[
							'block max-w-full truncate text-left text-sm',
							todo.done ? 'text-outline font-medium line-through' : 'text-on-surface font-semibold'
						]}
						title="Rename task"
						onclick={() => (mode = 'rename')}
					>
						{todo.title}
					</button>

					<div class="mt-2 flex flex-wrap items-center gap-2">
						<span
							class="bg-primary-fixed text-on-primary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
						>
							{projectTitle}
						</span>

						<span class="text-outline flex items-center gap-1 text-[11px]">
							<span class="material-symbols-outlined text-xs">
								{todo.done ? 'done_all' : 'schedule'}
							</span>

							{todo.done ? 'Completed' : 'Open'}
						</span>
					</div>
				</div>
			</div>

			<div
				class="flex shrink-0 items-center opacity-0 transition-opacity group-focus-within:opacity-100 group-hover:opacity-100"
			>
				<button
					type="button"
					class="text-on-surface-variant hover:bg-surface-container-high flex h-8 w-8 items-center justify-center rounded-lg"
					title="Rename task"
					aria-label={`Rename ${todo.title}`}
					onclick={() => (mode = 'rename')}
				>
					<span class="material-symbols-outlined text-base"> edit </span>
				</button>

				<button
					type="button"
					class="text-error hover:bg-error-container flex h-8 w-8 items-center justify-center rounded-lg"
					title="Delete task"
					aria-label={`Delete ${todo.title}`}
					onclick={() => (mode = 'delete')}
				>
					<span class="material-symbols-outlined text-base"> delete </span>
				</button>
			</div>
		</div>
	{/if}
</article>
