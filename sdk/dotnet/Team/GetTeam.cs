// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Team
{
    public static class GetTeam
    {
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("heroku:team/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("heroku:team/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        public readonly bool Default;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int MembershipLimit;
        public readonly string Name;
        public readonly bool ProvisionedLicenses;
        public readonly string Type;

        [OutputConstructor]
        private GetTeamResult(
            bool @default,

            string id,

            int membershipLimit,

            string name,

            bool provisionedLicenses,

            string type)
        {
            Default = @default;
            Id = id;
            MembershipLimit = membershipLimit;
            Name = name;
            ProvisionedLicenses = provisionedLicenses;
            Type = type;
        }
    }
}
