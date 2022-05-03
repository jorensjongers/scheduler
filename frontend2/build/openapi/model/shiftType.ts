/**
 * Scheduler API
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * Identifies a specific shift type. Includes the FREE shift type.
 */
export type ShiftType = 'JAEV' | 'JAWE' | 'JAHO' | 'SANW' | 'SAWE' | 'SAHO' | 'SAEV1' | 'SAEV2' | 'TPWE' | 'TPHO' | 'TPNF' | 'CALL' | 'FREE';

export const ShiftType = {
    Jaev: 'JAEV' as ShiftType,
    Jawe: 'JAWE' as ShiftType,
    Jaho: 'JAHO' as ShiftType,
    Sanw: 'SANW' as ShiftType,
    Sawe: 'SAWE' as ShiftType,
    Saho: 'SAHO' as ShiftType,
    Saev1: 'SAEV1' as ShiftType,
    Saev2: 'SAEV2' as ShiftType,
    Tpwe: 'TPWE' as ShiftType,
    Tpho: 'TPHO' as ShiftType,
    Tpnf: 'TPNF' as ShiftType,
    Call: 'CALL' as ShiftType,
    Free: 'FREE' as ShiftType
};

