// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a Heroku Addon Resource to an additional Heroku App.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * Addons can be imported using the unique Addon Attachment `id`, e.g.
 *
 * ```sh
 * $ pulumi import heroku:addon/attachment:Attachment foobar 01234567-89ab-cdef-0123-456789abcdef
 * ```
 */
export class Attachment extends pulumi.CustomResource {
    /**
     * Get an existing Attachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachmentState, opts?: pulumi.CustomResourceOptions): Attachment {
        return new Attachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'heroku:addon/attachment:Attachment';

    /**
     * Returns true if the given object is an instance of Attachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attachment.__pulumiType;
    }

    /**
     * The ID of the existing Heroku Addon to attach.
     */
    public readonly addonId!: pulumi.Output<string>;
    /**
     * Heroku app ID (do not use app name)
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * A friendly name for the Heroku Addon Attachment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
     */
    public readonly namespace!: pulumi.Output<string>;

    /**
     * Create a Attachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachmentArgs | AttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttachmentState | undefined;
            resourceInputs["addonId"] = state ? state.addonId : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if ((!args || args.addonId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addonId'");
            }
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            resourceInputs["addonId"] = args ? args.addonId : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Attachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    /**
     * The ID of the existing Heroku Addon to attach.
     */
    addonId?: pulumi.Input<string>;
    /**
     * Heroku app ID (do not use app name)
     */
    appId?: pulumi.Input<string>;
    /**
     * A friendly name for the Heroku Addon Attachment.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    /**
     * The ID of the existing Heroku Addon to attach.
     */
    addonId: pulumi.Input<string>;
    /**
     * Heroku app ID (do not use app name)
     */
    appId: pulumi.Input<string>;
    /**
     * A friendly name for the Heroku Addon Attachment.
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
     */
    namespace?: pulumi.Input<string>;
}
