// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a [Heroku App Webhook](https://devcenter.heroku.com/categories/app-webhooks).
 *
 * ## Example Usage
 *
 * ## Importing
 *
 * Existing webhooks can be imported using the combination of the application name or id, a colon, and the webhook name or id, e.g.
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:app/webhook:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * Heroku app ID (do not use app name)
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    public readonly authorization!: pulumi.Output<string | undefined>;
    /**
     * List of events to deliver to the webhook.
     */
    public readonly includes!: pulumi.Output<string[]>;
    /**
     * The webhook level (either `notify` or `sync`)
     */
    public readonly level!: pulumi.Output<string>;
    /**
     * Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    public readonly secret!: pulumi.Output<string | undefined>;
    /**
     * Optional plan configuration.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookArgs | WebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebhookState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["authorization"] = state ? state.authorization : undefined;
            resourceInputs["includes"] = state ? state.includes : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as WebhookArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.includes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'includes'");
            }
            if ((!args || args.level === undefined) && !opts.urn) {
                throw new Error("Missing required property 'level'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["authorization"] = args?.authorization ? pulumi.secret(args.authorization) : undefined;
            resourceInputs["includes"] = args ? args.includes : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["secret"] = args?.secret ? pulumi.secret(args.secret) : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authorization", "secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Webhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Webhook resources.
 */
export interface WebhookState {
    /**
     * Heroku app ID (do not use app name)
     */
    appId?: pulumi.Input<string>;
    /**
     * Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    authorization?: pulumi.Input<string>;
    /**
     * List of events to deliver to the webhook.
     */
    includes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The webhook level (either `notify` or `sync`)
     */
    level?: pulumi.Input<string>;
    /**
     * Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    secret?: pulumi.Input<string>;
    /**
     * Optional plan configuration.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * Heroku app ID (do not use app name)
     */
    appId: pulumi.Input<string>;
    /**
     * Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    authorization?: pulumi.Input<string>;
    /**
     * List of events to deliver to the webhook.
     */
    includes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The webhook level (either `notify` or `sync`)
     */
    level: pulumi.Input<string>;
    /**
     * Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
     */
    secret?: pulumi.Input<string>;
    /**
     * Optional plan configuration.
     */
    url: pulumi.Input<string>;
}
