// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.App
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// The most recent app release can be imported using the application name.
    /// 
    /// For example:
    /// 
    /// ```sh
    /// $ pulumi import heroku:app/release:Release foobar-release foobar
    /// ```
    /// </summary>
    [HerokuResourceType("heroku:app/release:Release")]
    public partial class Release : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// description of changes in this release
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// unique identifier of slug
        /// </summary>
        [Output("slugId")]
        public Output<string> SlugId { get; private set; } = null!;


        /// <summary>
        /// Create a Release resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Release(string name, ReleaseArgs args, CustomResourceOptions? options = null)
            : base("heroku:app/release:Release", name, args ?? new ReleaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Release(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
            : base("heroku:app/release:Release", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Release resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Release Get(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Release(name, id, state, options);
        }
    }

    public sealed class ReleaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// description of changes in this release
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// unique identifier of slug
        /// </summary>
        [Input("slugId", required: true)]
        public Input<string> SlugId { get; set; } = null!;

        public ReleaseArgs()
        {
        }
        public static new ReleaseArgs Empty => new ReleaseArgs();
    }

    public sealed class ReleaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// description of changes in this release
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// unique identifier of slug
        /// </summary>
        [Input("slugId")]
        public Input<string>? SlugId { get; set; }

        public ReleaseState()
        {
        }
        public static new ReleaseState Empty => new ReleaseState();
    }
}
