// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource is used to create and manage [User Features](https://devcenter.heroku.com/articles/heroku-beta-features) on Heroku.
 *
 * > **NOTE:** If this resource's HCL is removed from a `.tf` file, the behavior is to disable account feature
 * and remove resource from state.
 *
 * ## Available Features
 *
 * For a list of available features, use the [`heroku labs`](https://devcenter.heroku.com/articles/heroku-cli)
 * command to fetch them for the current authenticated user.
 *
 * The output will contain **User Features** that may be managed with this resource.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Existing account features can be imported using a combination of the account email (the email address tied to the Heroku API key)
 * and the feature name.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import heroku:account/feature:Feature example_metrics name@example.com:metrics-request-volume
 * ```
 */
export class Feature extends pulumi.CustomResource {
    /**
     * Get an existing Feature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FeatureState, opts?: pulumi.CustomResourceOptions): Feature {
        return new Feature(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:account/feature:Feature';

    /**
     * Returns true if the given object is an instance of Feature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Feature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Feature.__pulumiType;
    }

    /**
     * Description of account feature
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Enable or disable the account feature
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Name of the account feature
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * State of account feature
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Feature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeatureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FeatureArgs | FeatureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FeatureState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as FeatureArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Feature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Feature resources.
 */
export interface FeatureState {
    /**
     * Description of account feature
     */
    description?: pulumi.Input<string>;
    /**
     * Enable or disable the account feature
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the account feature
     */
    name?: pulumi.Input<string>;
    /**
     * State of account feature
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Feature resource.
 */
export interface FeatureArgs {
    /**
     * Enable or disable the account feature
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Name of the account feature
     */
    name?: pulumi.Input<string>;
}
