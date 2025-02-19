// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource for configuring review apps. Using this resource also enables review apps for a pipeline.
 *
 * > **IMPORTANT!**
 * This resource can only be used after you create a pipeline, and the pipeline has been connected to a Github repository.
 * Please visit this [help article](https://devcenter.heroku.com/articles/github-integration-review-apps#setup)
 * for more information.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * An Existing review app config using the combination of the pipeline UUID and the Github organization/repository
 *
 * separated by a colon.
 *
 * ```sh
 * $ pulumi import heroku:review/appConfig:AppConfig foobar afd193fb-7c5a-4d8f-afad-2388f4e6049d:heroku/homebrew-brew
 * ```
 */
export class AppConfig extends pulumi.CustomResource {
    /**
     * Get an existing AppConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppConfigState, opts?: pulumi.CustomResourceOptions): AppConfig {
        return new AppConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:review/appConfig:AppConfig';

    /**
     * Returns true if the given object is an instance of AppConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppConfig.__pulumiType;
    }

    /**
     * If true, this will trigger the creation of review apps when pull-requests
     * are opened in the repo. Defaults to `false`.
     */
    public readonly automaticReviewApps!: pulumi.Output<boolean | undefined>;
    /**
     * A unique prefix that will be used to create review app names.
     */
    public readonly baseName!: pulumi.Output<string>;
    /**
     * Provides a key/value pair to specify whether to use a common runtime or a private space.
     */
    public readonly deployTarget!: pulumi.Output<outputs.review.AppConfigDeployTarget>;
    /**
     * If `true`, this will trigger automatic deletion of review apps when they’re stale.
     * Defaults to `false`.
     */
    public readonly destroyStaleApps!: pulumi.Output<boolean | undefined>;
    /**
     * The fullName of the repository that you want to enable review-apps from.
     * Example `heroku/homebrew-brew`.
     */
    public readonly orgRepo!: pulumi.Output<string>;
    /**
     * The UUID of an existing pipeline.
     */
    public readonly pipelineId!: pulumi.Output<string>;
    public /*out*/ readonly repoId!: pulumi.Output<number>;
    /**
     * Destroy stale review apps automatically after these many days without any deploys.
     * Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
     */
    public readonly staleDays!: pulumi.Output<number>;
    /**
     * If true, review apps will only be created when CI passes. Defaults to `false`.
     */
    public readonly waitForCi!: pulumi.Output<boolean | undefined>;

    /**
     * Create a AppConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppConfigArgs | AppConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppConfigState | undefined;
            resourceInputs["automaticReviewApps"] = state ? state.automaticReviewApps : undefined;
            resourceInputs["baseName"] = state ? state.baseName : undefined;
            resourceInputs["deployTarget"] = state ? state.deployTarget : undefined;
            resourceInputs["destroyStaleApps"] = state ? state.destroyStaleApps : undefined;
            resourceInputs["orgRepo"] = state ? state.orgRepo : undefined;
            resourceInputs["pipelineId"] = state ? state.pipelineId : undefined;
            resourceInputs["repoId"] = state ? state.repoId : undefined;
            resourceInputs["staleDays"] = state ? state.staleDays : undefined;
            resourceInputs["waitForCi"] = state ? state.waitForCi : undefined;
        } else {
            const args = argsOrState as AppConfigArgs | undefined;
            if ((!args || args.deployTarget === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployTarget'");
            }
            if ((!args || args.orgRepo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgRepo'");
            }
            if ((!args || args.pipelineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineId'");
            }
            resourceInputs["automaticReviewApps"] = args ? args.automaticReviewApps : undefined;
            resourceInputs["baseName"] = args ? args.baseName : undefined;
            resourceInputs["deployTarget"] = args ? args.deployTarget : undefined;
            resourceInputs["destroyStaleApps"] = args ? args.destroyStaleApps : undefined;
            resourceInputs["orgRepo"] = args ? args.orgRepo : undefined;
            resourceInputs["pipelineId"] = args ? args.pipelineId : undefined;
            resourceInputs["staleDays"] = args ? args.staleDays : undefined;
            resourceInputs["waitForCi"] = args ? args.waitForCi : undefined;
            resourceInputs["repoId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppConfig resources.
 */
export interface AppConfigState {
    /**
     * If true, this will trigger the creation of review apps when pull-requests
     * are opened in the repo. Defaults to `false`.
     */
    automaticReviewApps?: pulumi.Input<boolean>;
    /**
     * A unique prefix that will be used to create review app names.
     */
    baseName?: pulumi.Input<string>;
    /**
     * Provides a key/value pair to specify whether to use a common runtime or a private space.
     */
    deployTarget?: pulumi.Input<inputs.review.AppConfigDeployTarget>;
    /**
     * If `true`, this will trigger automatic deletion of review apps when they’re stale.
     * Defaults to `false`.
     */
    destroyStaleApps?: pulumi.Input<boolean>;
    /**
     * The fullName of the repository that you want to enable review-apps from.
     * Example `heroku/homebrew-brew`.
     */
    orgRepo?: pulumi.Input<string>;
    /**
     * The UUID of an existing pipeline.
     */
    pipelineId?: pulumi.Input<string>;
    repoId?: pulumi.Input<number>;
    /**
     * Destroy stale review apps automatically after these many days without any deploys.
     * Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
     */
    staleDays?: pulumi.Input<number>;
    /**
     * If true, review apps will only be created when CI passes. Defaults to `false`.
     */
    waitForCi?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AppConfig resource.
 */
export interface AppConfigArgs {
    /**
     * If true, this will trigger the creation of review apps when pull-requests
     * are opened in the repo. Defaults to `false`.
     */
    automaticReviewApps?: pulumi.Input<boolean>;
    /**
     * A unique prefix that will be used to create review app names.
     */
    baseName?: pulumi.Input<string>;
    /**
     * Provides a key/value pair to specify whether to use a common runtime or a private space.
     */
    deployTarget: pulumi.Input<inputs.review.AppConfigDeployTarget>;
    /**
     * If `true`, this will trigger automatic deletion of review apps when they’re stale.
     * Defaults to `false`.
     */
    destroyStaleApps?: pulumi.Input<boolean>;
    /**
     * The fullName of the repository that you want to enable review-apps from.
     * Example `heroku/homebrew-brew`.
     */
    orgRepo: pulumi.Input<string>;
    /**
     * The UUID of an existing pipeline.
     */
    pipelineId: pulumi.Input<string>;
    /**
     * Destroy stale review apps automatically after these many days without any deploys.
     * Must be set with `destroyStaleApps` and value needs to be between `1` and `30` inclusive.
     */
    staleDays?: pulumi.Input<number>;
    /**
     * If true, review apps will only be created when CI passes. Defaults to `false`.
     */
    waitForCi?: pulumi.Input<boolean>;
}
