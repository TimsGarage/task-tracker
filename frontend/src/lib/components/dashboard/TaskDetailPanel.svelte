<script>
	import ConfirmAction from '../ui/ConfirmAction.svelte';

	let {
		todo,
		projects,
		onclose,
		onsave,
		ondelete,
		oncreatesubtask,
		onupdatesubtask,
		ondeletesubtask
	} = $props();

	let title = $state('');
	let done = $state(false);
	let projectId = $state(0);
	let dueDate = $state('');
	let priority = $state('medium');
	let tags = $state('');
	let notes = $state('');

	let saving = $state(false);
	let deleting = $state(false);
	let newSubtaskTitle = $state('');
	let addingSubtask = $state(false);
	let editingSubtaskId = $state(null);
	let subtaskDraft = $state('');

	let canSave = $derived(title.trim().length > 0 && !saving);

	let completedSubtasks = $derived(todo.subtasks?.filter((subtask) => subtask.done).length ?? 0);

	$effect(() => {
		todo.id;

		title = todo.title;
		done = todo.done;
		projectId = todo.projectId;
		dueDate = toLocalDateTime(todo.dueDate);
		priority = todo.priority || 'medium';
		tags = (todo.tags ?? []).join(', ');
		notes = todo.notes ?? '';
	});

	function toLocalDateTime(value) {
		if (!value) return '';

		const date = new Date(value);
		const offset = date.getTimezoneOffset() * 60_000;

		return new Date(date.getTime() - offset).toISOString().slice(0, 16);
	}

	function toApiDate(value) {
		return value ? new Date(value).toISOString() : null;
	}

	function parseTags(value) {
		return [
			...new Set(
				value
					.split(',')
					.map((tag) => tag.trim().replace(/^#/, ''))
					.filter(Boolean)
			)
		];
	}

	async function save() {
		if (!canSave) return;

		saving = true;

		try {
			await onsave(todo.id, {
				title: title.trim(),
				done,
				projectId: Number(projectId),
				dueDate: toApiDate(dueDate),
				priority,
				notes,
				tags: parseTags(tags)
			});
		} catch {
			// Handled by the app-level banner.
		} finally {
			saving = false;
		}
	}

	async function addSubtask(event) {
		event.preventDefault();

		const value = newSubtaskTitle.trim();
		if (!value || addingSubtask) return;

		addingSubtask = true;

		try {
			await oncreatesubtask(todo.id, value);
			newSubtaskTitle = '';
		} catch {
			// Handled by the app-level banner.
		} finally {
			addingSubtask = false;
		}
	}

	function beginSubtaskRename(subtask) {
		editingSubtaskId = subtask.id;
		subtaskDraft = subtask.title;
	}

	async function saveSubtaskRename(subtask) {
		const value = subtaskDraft.trim();
		if (!value) return;

		try {
			await onupdatesubtask(subtask.id, {
				title: value
			});

			editingSubtaskId = null;
			subtaskDraft = '';
		} catch {
			// Handled by the app-level banner.
		}
	}

	async function deleteTask() {
		await ondelete(todo.id);
		deleting = false;
	}
</script>

<aside
	class="bg-surface-container-lowest sticky top-24 flex flex-col gap-6 rounded-xl p-6 shadow-sm xl:col-span-5"
>
	<div class="flex items-center justify-between">
		<div class="flex items-center gap-2">
			<span class="bg-primary h-2.5 w-2.5 rounded-full"></span>

			<span class="text-secondary text-xs font-semibold tracking-wider uppercase">
				Task details
			</span>
		</div>

		<button
			type="button"
			class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg"
			title="Close panel"
			onclick={onclose}
		>
			<span class="material-symbols-outlined text-lg"> close </span>
		</button>
	</div>

	<div class="flex items-start gap-3">
		<button
			type="button"
			class={[
				'mt-1 flex h-6 w-6 shrink-0 items-center justify-center rounded-full transition-colors',
				done ? 'bg-primary text-on-primary' : 'border-primary/70 border-2'
			]}
			onclick={() => (done = !done)}
			aria-label={done ? 'Mark as incomplete' : 'Mark as complete'}
		>
			{#if done}
				<span class="material-symbols-outlined text-sm"> check </span>
			{/if}
		</button>

		<textarea
			bind:value={title}
			rows="2"
			class="font-headline text-on-surface focus:bg-surface-container-low min-h-16 flex-1 resize-none rounded-lg bg-transparent px-2 py-1 text-lg leading-snug font-semibold focus:outline-none"
			aria-label="Task title"></textarea>
	</div>

	<div class="bg-surface-container-low grid grid-cols-1 gap-3 rounded-xl p-3.5 sm:grid-cols-2">
		<label class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Due date
			</span>

			<input
				bind:value={dueDate}
				type="datetime-local"
				class="bg-surface-container-lowest text-on-surface rounded-lg border-0 px-3 py-2 text-xs"
			/>
		</label>

		<label class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Priority
			</span>

			<select
				bind:value={priority}
				class="bg-surface-container-lowest text-on-surface rounded-lg border-0 px-3 py-2 text-xs"
			>
				<option value="low">Low priority</option>
				<option value="medium">Medium priority</option>
				<option value="high">High priority</option>
			</select>
		</label>

		<label class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase">
				Project
			</span>

			<select
				bind:value={projectId}
				class="bg-surface-container-lowest text-on-surface rounded-lg border-0 px-3 py-2 text-xs"
			>
				{#each projects as project (project.id)}
					<option value={project.id}>
						{project.title}
					</option>
				{/each}
			</select>
		</label>

		<label class="flex flex-col gap-1">
			<span class="text-secondary text-[11px] font-semibold tracking-wider uppercase"> Tags </span>

			<input
				bind:value={tags}
				class="bg-surface-container-lowest text-on-surface rounded-lg border-0 px-3 py-2 text-xs"
				placeholder="garden, urgent"
			/>
		</label>
	</div>

	<label class="space-y-2">
		<span class="text-secondary text-xs font-semibold tracking-wider uppercase"> Field notes </span>

		<textarea
			bind:value={notes}
			rows="4"
			class="bg-surface-container-low text-on-surface placeholder:text-outline w-full resize-y rounded-xl border-0 p-3.5 text-xs leading-relaxed focus:outline-none"
			placeholder="Add observations or reminders..."></textarea>
	</label>

	<section class="space-y-3">
		<div class="flex items-center justify-between">
			<span class="text-secondary text-xs font-semibold tracking-wider uppercase">
				Subtasks ({completedSubtasks} of {todo.subtasks?.length ?? 0})
			</span>
		</div>

		<form class="flex gap-2" onsubmit={addSubtask}>
			<input
				bind:value={newSubtaskTitle}
				class="bg-surface-container-low text-on-surface min-w-0 flex-1 rounded-lg border-0 px-3 py-2 text-xs"
				placeholder="New subtask"
				disabled={addingSubtask}
			/>

			<button
				type="submit"
				class="bg-primary text-on-primary rounded-lg px-3 py-2 text-xs disabled:opacity-50"
				disabled={!newSubtaskTitle.trim() || addingSubtask}
			>
				Add
			</button>
		</form>

		<div class="space-y-2">
			{#each todo.subtasks ?? [] as subtask (subtask.id)}
				<div class="bg-surface-container-low group flex items-center gap-2 rounded-lg p-2">
					<button
						type="button"
						class={[
							'flex h-4 w-4 shrink-0 items-center justify-center rounded-full',
							subtask.done ? 'bg-primary text-on-primary' : 'border-outline border'
						]}
						onclick={() =>
							onupdatesubtask(subtask.id, {
								done: !subtask.done
							})}
					>
						{#if subtask.done}
							<span class="material-symbols-outlined text-[10px]"> check </span>
						{/if}
					</button>

					{#if editingSubtaskId === subtask.id}
						<input
							bind:value={subtaskDraft}
							class="bg-surface-container-lowest min-w-0 flex-1 rounded-md border-0 px-2 py-1 text-xs"
							onkeydown={(event) => {
								if (event.key === 'Enter') {
									event.preventDefault();
									saveSubtaskRename(subtask);
								}

								if (event.key === 'Escape') {
									editingSubtaskId = null;
								}
							}}
						/>

						<button
							type="button"
							class="text-primary text-xs"
							onclick={() => saveSubtaskRename(subtask)}
						>
							Save
						</button>
					{:else}
						<button
							type="button"
							class={[
								'min-w-0 flex-1 truncate text-left text-xs',
								subtask.done ? 'text-outline line-through' : 'text-on-surface'
							]}
							onclick={() => beginSubtaskRename(subtask)}
						>
							{subtask.title}
						</button>

						<button
							type="button"
							class="text-error opacity-0 transition-opacity group-hover:opacity-100"
							aria-label={`Delete ${subtask.title}`}
							onclick={() => ondeletesubtask(subtask.id)}
						>
							<span class="material-symbols-outlined text-sm"> delete </span>
						</button>
					{/if}
				</div>
			{/each}
		</div>
	</section>

	{#if deleting}
		<ConfirmAction
			message={`Delete "${todo.title}"?`}
			confirmLabel="Delete task"
			onconfirm={deleteTask}
			oncancel={() => (deleting = false)}
		/>
	{:else}
		<div class="flex items-center justify-between gap-3">
			<button
				type="button"
				class="text-error flex items-center gap-1 text-xs font-medium"
				onclick={() => (deleting = true)}
			>
				<span class="material-symbols-outlined text-sm"> delete_outline </span>

				Delete task
			</button>

			<button
				type="button"
				class="bg-primary text-on-primary rounded-xl px-4 py-2 text-xs font-medium shadow-sm disabled:opacity-50"
				disabled={!canSave}
				onclick={save}
			>
				{saving ? 'Saving…' : 'Save changes'}
			</button>
		</div>
	{/if}
</aside>
