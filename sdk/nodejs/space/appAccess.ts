// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for managing permissions for the entire Private Space. Members with the admin role will
 * always have full permissions in the Private Space, so using this resource on an admin will have no effect.
 * The provided email must already be a member of the Heroku Team. Currently, the only supported permission is `createApps`.
 *
 * ## Example Usage
 *
 * ## Importing
 *
 * Existing permissions can be imported using the combination of the Private Space name, a colon, and the member email.
 *
 * For example:
 */
export class AppAccess extends pulumi.CustomResource {
    /**
     * Get an existing AppAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppAccessState, opts?: pulumi.CustomResourceOptions): AppAccess {
        return new AppAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:space/appAccess:AppAccess';

    /**
     * Returns true if the given object is an instance of AppAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppAccess.__pulumiType;
    }

    /**
     * The email of the existing Heroku Team member.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The permissions to grant the team member for the Private Space.
     * Currently `createApps` is the only supported permission. If not provided the member will have no permissions to the space.
     * Members with admin role will always have `createApps` permissions, which cannot be removed.
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * The ID of the Private Space.
     */
    public readonly space!: pulumi.Output<string>;

    /**
     * Create a AppAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppAccessArgs | AppAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppAccessState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["space"] = state ? state.space : undefined;
        } else {
            const args = argsOrState as AppAccessArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.space === undefined) && !opts.urn) {
                throw new Error("Missing required property 'space'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["space"] = args ? args.space : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppAccess resources.
 */
export interface AppAccessState {
    /**
     * The email of the existing Heroku Team member.
     */
    email?: pulumi.Input<string>;
    /**
     * The permissions to grant the team member for the Private Space.
     * Currently `createApps` is the only supported permission. If not provided the member will have no permissions to the space.
     * Members with admin role will always have `createApps` permissions, which cannot be removed.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Private Space.
     */
    space?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AppAccess resource.
 */
export interface AppAccessArgs {
    /**
     * The email of the existing Heroku Team member.
     */
    email: pulumi.Input<string>;
    /**
     * The permissions to grant the team member for the Private Space.
     * Currently `createApps` is the only supported permission. If not provided the member will have no permissions to the space.
     * Members with admin role will always have `createApps` permissions, which cannot be removed.
     */
    permissions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the Private Space.
     */
    space: pulumi.Input<string>;
}
