// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Space.Inputs
{

    public sealed class InboundRulesetRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public InboundRulesetRuleArgs()
        {
        }
        public static new InboundRulesetRuleArgs Empty => new InboundRulesetRuleArgs();
    }
}