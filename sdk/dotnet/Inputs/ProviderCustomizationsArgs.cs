// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Inputs
{

    public sealed class ProviderCustomizationsArgs : global::Pulumi.ResourceArgs
    {
        [Input("setAppAllConfigVarsInState")]
        public Input<bool>? SetAppAllConfigVarsInState { get; set; }

        public ProviderCustomizationsArgs()
        {
        }
        public static new ProviderCustomizationsArgs Empty => new ProviderCustomizationsArgs();
    }
}
