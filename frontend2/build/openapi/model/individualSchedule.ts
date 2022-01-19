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
import { Assignment } from './assignment';


/**
 * Represents the work schedule of an individual assistant.
 */
export interface IndividualSchedule { 
    /**
     * The identification number of the assistant for which this is an individual schedule.
     */
    assistant_id: number;
    /**
     * The workload of this inidividual schedule. Used when calculating the fairness score.
     */
    workload: number;
    /**
     * Contains all the individual assignments of this individual schedule.
     */
    assignments: Array<Assignment>;
}

