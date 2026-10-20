import { apiRequest } from './client.js';

/**
 * @returns {Promise<import("../domain/types.js").Project[]>}
 */
export function getProjects() {
	return apiRequest('/projects');
}

/**
 * @param {string} title
 * @returns {Promise<import("../domain/types.js").Project>}
 */
export function createProject(title) {
	return apiRequest('/projects', {
		method: 'POST',
		body: JSON.stringify({ title })
	});
}

/**
 * @param {number} projectId
 * @param {string} title
 * @returns {Promise<import("../domain/types.js").Project>}
 */
export function renameProject(projectId, title) {
	return apiRequest(`/projects/${projectId}`, {
		method: 'PATCH',
		body: JSON.stringify({ title })
	});
}

/**
 * @param {number} projectId
 * @returns {Promise<null>}
 */
export function deleteProject(projectId) {
	return apiRequest(`/projects/${projectId}`, {
		method: 'DELETE'
	});
}
