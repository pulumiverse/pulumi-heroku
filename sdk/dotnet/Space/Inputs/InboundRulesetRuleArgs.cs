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
        /// <summary>
        /// The action to apply this rule to. Must be one of `allow` or `deny`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// A CIDR block source for the rule.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public InboundRulesetRuleArgs()
        {
        }
        public static new InboundRulesetRuleArgs Empty => new InboundRulesetRuleArgs();
    }
}
