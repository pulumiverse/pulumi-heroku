// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the heroku package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'heroku';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    public readonly apiKey!: pulumi.Output<string | undefined>;
    public readonly email!: pulumi.Output<string | undefined>;
    public readonly headers!: pulumi.Output<string | undefined>;
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["apiKey"] = args ? args.apiKey : undefined;
            resourceInputs["customizations"] = pulumi.output(args ? args.customizations : undefined).apply(JSON.stringify);
            resourceInputs["delays"] = pulumi.output(args ? args.delays : undefined).apply(JSON.stringify);
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["headers"] = args ? args.headers : undefined;
            resourceInputs["timeouts"] = pulumi.output(args ? args.timeouts : undefined).apply(JSON.stringify);
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    apiKey?: pulumi.Input<string>;
    customizations?: pulumi.Input<inputs.ProviderCustomizations>;
    delays?: pulumi.Input<inputs.ProviderDelays>;
    email?: pulumi.Input<string>;
    headers?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.ProviderTimeouts>;
    url?: pulumi.Input<string>;
}
