import { createProject, deleteProject, getProjects, renameProject } from '../api/projects.js';

import { createTodo, deleteTodo, updateTodo } from '../api/todos.js';

export class TodoAppState {
	projects = $state([]);
	selectedProjectId = $state(null);
	error = $state('');
	loading = $state(true);
	composerFocusToken = $state(0);

	get selectedProject() {
		return this.projects.find((project) => project.id === this.selectedProjectId) ?? null;
	}

	get totalTodoCount() {
		return this.projects.reduce((total, project) => total + (project.todos?.length ?? 0), 0);
	}

	get completedTodoCount() {
		return this.projects.reduce(
			(total, project) => total + (project.todos?.filter((todo) => todo.done).length ?? 0),
			0
		);
	}

	get completionPercentage() {
		if (this.totalTodoCount === 0) {
			return 0;
		}

		return Math.round((this.completedTodoCount / this.totalTodoCount) * 100);
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

		const preferredProjectExists = projects.some((project) => project.id === preferredProjectId);

		this.selectedProjectId = preferredProjectExists
			? preferredProjectId
			: (projects[0]?.id ?? null);
	}

	async load() {
		this.loading = true;

		try {
			await this.perform(() => this.refresh(null));
		} catch {
			// The error is exposed through this.error.
		} finally {
			this.loading = false;
		}
	}

	selectProject(projectId) {
		this.selectedProjectId = projectId;
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
			await createTodo(projectId, title);
			await this.refresh(projectId);
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

	async removeTodo(todoId) {
		const projectId = this.selectedProjectId;

		return this.perform(async () => {
			await deleteTodo(todoId);
			await this.refresh(projectId);
		});
	}
}
