// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Build.Inputs
{

    public sealed class BuildSourceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("checksum")]
        public Input<string>? Checksum { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildSourceGetArgs()
        {
        }
        public static new BuildSourceGetArgs Empty => new BuildSourceGetArgs();
    }
}
