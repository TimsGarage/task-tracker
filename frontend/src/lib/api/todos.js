import { apiRequest } from './client.js';

/**
 * @param {number} projectId
 * @param {string} title
 */
export function createTodo(projectId, title) {
	return apiRequest(`/projects/${projectId}/todos`, {
		method: 'POST',
		body: JSON.stringify({ title })
	});
}

/**
 * @param {number} todoId
 * @param {{
 *   title?: string,
 *   done?: boolean,
 *   projectId?: number,
 *   dueDate?: string | null,
 *   priority?: "low" | "medium" | "high",
 *   notes?: string,
 *   tags?: string[]
 * }} changes
 */
export function updateTodo(todoId, changes) {
	return apiRequest(`/todos/${todoId}`, {
		method: 'PATCH',
		body: JSON.stringify(changes)
	});
}

/**
 * @param {number} todoId
 */
export function deleteTodo(todoId) {
	return apiRequest(`/todos/${todoId}`, {
		method: 'DELETE'
	});
}

/**
 * @param {number} todoId
 * @param {string} title
 */
export function createSubtask(todoId, title) {
	return apiRequest(`/todos/${todoId}/subtasks`, {
		method: 'POST',
		body: JSON.stringify({ title })
	});
}

/**
 * @param {number} subtaskId
 * @param {{ title?: string, done?: boolean }} changes
 */
export function updateSubtask(subtaskId, changes) {
	return apiRequest(`/subtasks/${subtaskId}`, {
		method: 'PATCH',
		body: JSON.stringify(changes)
	});
}

/**
 * @param {number} subtaskId
 */
export function deleteSubtask(subtaskId) {
	return apiRequest(`/subtasks/${subtaskId}`, {
		method: 'DELETE'
	});
}
