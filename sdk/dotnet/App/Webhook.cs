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
    /// Provides a [Heroku App Webhook](https://devcenter.heroku.com/categories/app-webhooks).
    /// 
    /// ## Example Usage
    /// 
    /// ## Importing
    /// 
    /// Existing webhooks can be imported using the combination of the application name or id, a colon, and the webhook name or id, e.g.
    /// </summary>
    [HerokuResourceType("heroku:app/webhook:Webhook")]
    public partial class Webhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        [Output("authorization")]
        public Output<string?> Authorization { get; private set; } = null!;

        /// <summary>
        /// List of events to deliver to the webhook.
        /// </summary>
        [Output("includes")]
        public Output<ImmutableArray<string>> Includes { get; private set; } = null!;

        /// <summary>
        /// The webhook level (either `notify` or `sync`)
        /// </summary>
        [Output("level")]
        public Output<string> Level { get; private set; } = null!;

        /// <summary>
        /// Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        [Output("secret")]
        public Output<string?> Secret { get; private set; } = null!;

        /// <summary>
        /// Optional plan configuration.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Webhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webhook(string name, WebhookArgs args, CustomResourceOptions? options = null)
            : base("heroku:app/webhook:Webhook", name, args ?? new WebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webhook(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
            : base("heroku:app/webhook:Webhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "authorization",
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Webhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webhook Get(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new Webhook(name, id, state, options);
        }
    }

    public sealed class WebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("authorization")]
        private Input<string>? _authorization;

        /// <summary>
        /// Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        public Input<string>? Authorization
        {
            get => _authorization;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authorization = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("includes", required: true)]
        private InputList<string>? _includes;

        /// <summary>
        /// List of events to deliver to the webhook.
        /// </summary>
        public InputList<string> Includes
        {
            get => _includes ?? (_includes = new InputList<string>());
            set => _includes = value;
        }

        /// <summary>
        /// The webhook level (either `notify` or `sync`)
        /// </summary>
        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Optional plan configuration.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public WebhookArgs()
        {
        }
        public static new WebhookArgs Empty => new WebhookArgs();
    }

    public sealed class WebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("authorization")]
        private Input<string>? _authorization;

        /// <summary>
        /// Values used in `Authorization` header. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        public Input<string>? Authorization
        {
            get => _authorization;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authorization = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("includes")]
        private InputList<string>? _includes;

        /// <summary>
        /// List of events to deliver to the webhook.
        /// </summary>
        public InputList<string> Includes
        {
            get => _includes ?? (_includes = new InputList<string>());
            set => _includes = value;
        }

        /// <summary>
        /// The webhook level (either `notify` or `sync`)
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// Value used to sign webhook payloads. Once set, this value cannot be fetched from the Heroku API, but it can be updated.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Optional plan configuration.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebhookState()
        {
        }
        public static new WebhookState Empty => new WebhookState();
    }
}
