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
/**
 * 
 * @export
 * @interface ControllersProvince
 */
export interface ControllersProvince {
    /**
     * 
     * @type {number}
     * @memberof ControllersProvince
     */
    continent?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProvince
     */
    country?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProvince
     */
    district?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersProvince
     */
    postal?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProvince
     */
    province?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersProvince
     */
    region?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProvince
     */
    subdistrict?: string;
}

export function ControllersProvinceFromJSON(json: any): ControllersProvince {
    return ControllersProvinceFromJSONTyped(json, false);
}

export function ControllersProvinceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersProvince {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'continent': !exists(json, 'continent') ? undefined : json['continent'],
        'country': !exists(json, 'country') ? undefined : json['country'],
        'district': !exists(json, 'district') ? undefined : json['district'],
        'postal': !exists(json, 'postal') ? undefined : json['postal'],
        'province': !exists(json, 'province') ? undefined : json['province'],
        'region': !exists(json, 'region') ? undefined : json['region'],
        'subdistrict': !exists(json, 'subdistrict') ? undefined : json['subdistrict'],
    };
}

export function ControllersProvinceToJSON(value?: ControllersProvince | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'continent': value.continent,
        'country': value.country,
        'district': value.district,
        'postal': value.postal,
        'province': value.province,
        'region': value.region,
        'subdistrict': value.subdistrict,
    };
}

