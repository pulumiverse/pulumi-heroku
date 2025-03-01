// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A [Heroku Team Member](https://devcenter.heroku.com/articles/platform-api-reference#team-member) receives access to everything owned by the Team.
 *
 * To create a Heroku Team, use the [New Team](https://dashboard.heroku.com/teams/new) feature of Heroku Dashboard. For Heroku Enterprise accounts, new Teams may be created within the account by users with the right permissions.
 *
 * A Heroku "team" was originally called an "organization", and that is still the identifier used elsewhere in this provider. For `heroku.app.App` & `heroku.space.Space` resources, set the Heroku Team name as the "organization".
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Team members can be imported using the combination of the team application name, a colon, and the member's email address.
 *
 * ```sh
 * $ pulumi import heroku:team/member:Member foobar-member my-team-foobar:some-user@example.com
 * ```
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberState, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:team/member:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    /**
     * Email address of the member
     */
    public readonly email!: pulumi.Output<string>;
    public readonly federated!: pulumi.Output<boolean | undefined>;
    /**
     * The role to assign the member. See [the API docs](https://devcenter.heroku.com/articles/platform-api-reference#team-member) for available options.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The name of the Heroku Team.
     */
    public readonly team!: pulumi.Output<string>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberArgs | MemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MemberState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["federated"] = state ? state.federated : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["team"] = state ? state.team : undefined;
        } else {
            const args = argsOrState as MemberArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.team === undefined) && !opts.urn) {
                throw new Error("Missing required property 'team'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["federated"] = args ? args.federated : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["team"] = args ? args.team : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Member.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Member resources.
 */
export interface MemberState {
    /**
     * Email address of the member
     */
    email?: pulumi.Input<string>;
    federated?: pulumi.Input<boolean>;
    /**
     * The role to assign the member. See [the API docs](https://devcenter.heroku.com/articles/platform-api-reference#team-member) for available options.
     */
    role?: pulumi.Input<string>;
    /**
     * The name of the Heroku Team.
     */
    team?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    /**
     * Email address of the member
     */
    email: pulumi.Input<string>;
    federated?: pulumi.Input<boolean>;
    /**
     * The role to assign the member. See [the API docs](https://devcenter.heroku.com/articles/platform-api-reference#team-member) for available options.
     */
    role: pulumi.Input<string>;
    /**
     * The name of the Heroku Team.
     */
    team: pulumi.Input<string>;
}
