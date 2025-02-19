// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Addon
{
    /// <summary>
    /// Attaches a Heroku Addon Resource to an additional Heroku App.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Addons can be imported using the unique Addon Attachment `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import heroku:addon/attachment:Attachment foobar 01234567-89ab-cdef-0123-456789abcdef
    /// ```
    /// </summary>
    [HerokuResourceType("heroku:addon/attachment:Attachment")]
    public partial class Attachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the existing Heroku Addon to attach.
        /// </summary>
        [Output("addonId")]
        public Output<string> AddonId { get; private set; } = null!;

        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// A friendly name for the Heroku Addon Attachment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attachment(string name, AttachmentArgs args, CustomResourceOptions? options = null)
            : base("heroku:addon/attachment:Attachment", name, args ?? new AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attachment(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("heroku:addon/attachment:Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attachment Get(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Attachment(name, id, state, options);
        }
    }

    public sealed class AttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the existing Heroku Addon to attach.
        /// </summary>
        [Input("addonId", required: true)]
        public Input<string> AddonId { get; set; } = null!;

        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// A friendly name for the Heroku Addon Attachment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public AttachmentArgs()
        {
        }
        public static new AttachmentArgs Empty => new AttachmentArgs();
    }

    public sealed class AttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the existing Heroku Addon to attach.
        /// </summary>
        [Input("addonId")]
        public Input<string>? AddonId { get; set; }

        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// A friendly name for the Heroku Addon Attachment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace value for the Heroku Addon Attachment. This can be used to configure the behaviour of the attachment. See [Heroku Platform API Reference](https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment-create)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public AttachmentState()
        {
        }
        public static new AttachmentState Empty => new AttachmentState();
    }
}
