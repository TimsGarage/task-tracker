import { createProject, deleteProject, getProjects, renameProject } from '../api/projects.js';

import {
	createSubtask,
	createTodo,
	deleteSubtask,
	deleteTodo,
	updateSubtask,
	updateTodo
} from '../api/todos.js';

export class TodoAppState {
	projects = $state([]);
	selectedProjectId = $state(null);
	selectedTodoId = $state(null);
	error = $state('');
	loading = $state(true);
	composerFocusToken = $state(0);

	get selectedProject() {
		return this.projects.find((project) => project.id === this.selectedProjectId) ?? null;
	}

	get selectedTodo() {
		if (this.selectedTodoId === null) {
			return null;
		}

		for (const project of this.projects) {
			const todo = project.todos?.find((item) => item.id === this.selectedTodoId);

			if (todo) {
				return todo;
			}
		}

		return null;
	}

	async perform(action) {
		this.error = '';

		try {
			return await action();
		} catch (error) {
			this.error = error instanceof Error ? error.message : 'An unexpected error occurred';

			throw error;
		}
	}

	async refresh(preferredProjectId = this.selectedProjectId) {
		const projects = await getProjects();

		this.projects = projects;

		const projectExists = projects.some((project) => project.id === preferredProjectId);

		this.selectedProjectId = projectExists ? preferredProjectId : (projects[0]?.id ?? null);

		if (this.selectedTodoId !== null && !this.selectedTodo) {
			this.selectedTodoId = null;
		}
	}

	async load() {
		this.loading = true;

		try {
			await this.perform(() => this.refresh(null));
		} catch {
			// The error is available through this.error.
		} finally {
			this.loading = false;
		}
	}

	selectProject(projectId) {
		this.selectedProjectId = projectId;
		this.selectedTodoId = null;
	}

	selectTodo(todoId) {
		this.selectedTodoId = todoId;
	}

	closeTodo() {
		this.selectedTodoId = null;
	}

	focusTodoComposer() {
		this.composerFocusToken += 1;
	}

	dismissError() {
		this.error = '';
	}

	async addProject(title) {
		return this.perform(async () => {
			const project = await createProject(title);
			await this.refresh(project.id);
		});
	}

	async editProject(projectId, title) {
		return this.perform(async () => {
			await renameProject(projectId, title);
			await this.refresh(projectId);
		});
	}

	async removeProject(projectId) {
		return this.perform(async () => {
			await deleteProject(projectId);
			await this.refresh();
		});
	}

	async addTodo(title) {
		if (!this.selectedProject) {
			return;
		}

		const projectId = this.selectedProject.id;

		return this.perform(async () => {
			const todo = await createTodo(projectId, title);
			await this.refresh(projectId);
			this.selectedTodoId = todo.id;
		});
	}

	async toggleTodo(todoId, done) {
		const projectId = this.selectedProjectId;

		return this.perform(async () => {
			await updateTodo(todoId, { done });
			await this.refresh(projectId);
		});
	}

	async editTodo(todoId, title) {
		const projectId = this.selectedProjectId;

		return this.perform(async () => {
			await updateTodo(todoId, { title });
			await this.refresh(projectId);
		});
	}

	async saveTodo(todoId, changes) {
		return this.perform(async () => {
			const updated = await updateTodo(todoId, changes);

			await this.refresh(updated.projectId);

			this.selectedProjectId = updated.projectId;
			this.selectedTodoId = updated.id;
		});
	}

	async removeTodo(todoId) {
		const projectId = this.selectedProjectId;

		return this.perform(async () => {
			await deleteTodo(todoId);
			this.selectedTodoId = null;
			await this.refresh(projectId);
		});
	}

	async addSubtask(todoId, title) {
		const projectId = this.selectedProjectId;

		return this.perform(async () => {
			await createSubtask(todoId, title);
			await this.refresh(projectId);
			this.selectedTodoId = todoId;
		});
	}

	async editSubtask(subtaskId, changes) {
		const projectId = this.selectedProjectId;
		const todoId = this.selectedTodoId;

		return this.perform(async () => {
			await updateSubtask(subtaskId, changes);
			await this.refresh(projectId);
			this.selectedTodoId = todoId;
		});
	}

	async removeSubtask(subtaskId) {
		const projectId = this.selectedProjectId;
		const todoId = this.selectedTodoId;

		return this.perform(async () => {
			await deleteSubtask(subtaskId);
			await this.refresh(projectId);
			this.selectedTodoId = todoId;
		});
	}
}
