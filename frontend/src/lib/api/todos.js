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
 * @param {{ title?: string, done?: boolean }} changes
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
