/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntFacultyEdges,
    EntFacultyEdgesFromJSON,
    EntFacultyEdgesFromJSONTyped,
    EntFacultyEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFaculty
 */
export interface EntFaculty {
    /**
     * 
     * @type {EntFacultyEdges}
     * @memberof EntFaculty
     */
    edges?: EntFacultyEdges;
    /**
     * Faculty holds the value of the "faculty" field.
     * @type {string}
     * @memberof EntFaculty
     */
    faculty?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntFaculty
     */
    id?: number;
}

export function EntFacultyFromJSON(json: any): EntFaculty {
    return EntFacultyFromJSONTyped(json, false);
}

export function EntFacultyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFaculty {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntFacultyEdgesFromJSON(json['edges']),
        'faculty': !exists(json, 'faculty') ? undefined : json['faculty'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFacultyToJSON(value?: EntFaculty | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntFacultyEdgesToJSON(value.edges),
        'faculty': value.faculty,
        'id': value.id,
    };
}


