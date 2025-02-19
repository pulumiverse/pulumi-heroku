// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * The most recent app release can be imported using the application name.
 *
 * For example:
 *
 * ```sh
 * $ pulumi import heroku:app/release:Release foobar-release foobar
 * ```
 */
export class Release extends pulumi.CustomResource {
    /**
     * Get an existing Release resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReleaseState, opts?: pulumi.CustomResourceOptions): Release {
        return new Release(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:app/release:Release';

    /**
     * Returns true if the given object is an instance of Release.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Release {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Release.__pulumiType;
    }

    /**
     * Heroku app ID (do not use app name)
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * description of changes in this release
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * unique identifier of slug
     */
    public readonly slugId!: pulumi.Output<string>;

    /**
     * Create a Release resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReleaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReleaseArgs | ReleaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReleaseState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["slugId"] = state ? state.slugId : undefined;
        } else {
            const args = argsOrState as ReleaseArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.slugId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'slugId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["slugId"] = args ? args.slugId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Release.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Release resources.
 */
export interface ReleaseState {
    /**
     * Heroku app ID (do not use app name)
     */
    appId?: pulumi.Input<string>;
    /**
     * description of changes in this release
     */
    description?: pulumi.Input<string>;
    /**
     * unique identifier of slug
     */
    slugId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Release resource.
 */
export interface ReleaseArgs {
    /**
     * Heroku app ID (do not use app name)
     */
    appId: pulumi.Input<string>;
    /**
     * description of changes in this release
     */
    description?: pulumi.Input<string>;
    /**
     * unique identifier of slug
     */
    slugId: pulumi.Input<string>;
}
